package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/genie-service/models"
	utils "github.com/genie-service/utils"
	"github.com/genie-service/utils/environment"
)

var client *ethclient.Client
var logTransferSig []byte
var logTransferSigHash common.Hash

func init() {
	wssUrl := environment.GetEnvVariable("ETHEREUM_NODE_WSS")

	wsClient, dialErr := ethclient.Dial(wssUrl)

	if dialErr != nil {
		panic(dialErr)
	}

	client = wsClient

	logTransferSig = []byte("Transfer(address,address,uint256)")
	logTransferSigHash = crypto.Keccak256Hash(logTransferSig)
}

func SubscribeToContractEvents(address string) {
	fmt.Println("Initiating WS to consume event for ", address, "\n")

	contractAddress := common.HexToAddress(address)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	subscription, subscriptionErr := client.SubscribeFilterLogs(
		context.Background(),
		query,
		logs,
	)

	if subscriptionErr != nil {
		panic(subscriptionErr)
	}

	for {
		select {
		case err := <-subscription.Err():
			log.Fatal(err)
		case eventLog := <-logs:
			// Limit parsing to "Transfer" Event Logs
			if strings.EqualFold(eventLog.Topics[0].Hex(), logTransferSigHash.Hex()) {
				fmt.Println("Processing Transfer Event from ", address)
				// Wait 10 seconds before initiating rarible API call to let their systems update
				// Just chose an arbritrary value right now but would want to read more on the SLA
				time.Sleep(10 * time.Second)

				// Update information for all assets in the collection since we can't decode indexed
				// attributes inside eventLog.Topics
				go ParseContractData(address)

				// parseContractLog(eventLog, address)
			}
		}
	}
}

func parseContractLog(log types.Log, address string) {
	// Get ABI for given address.
	contractAbi, err := abi.JSON(strings.NewReader(utils.Abis[address]))

	if err != nil {
		panic(err)
	}

	event := struct {
		From    common.Address
		To      common.Address
		TokenId *big.Int
	}{}

	transferEvents, unpackErr := contractAbi.Unpack("Transfer", log.Data)

	if unpackErr != nil {
		panic(unpackErr)
	}
	fmt.Println("data: ", log.Data)
	fmt.Println("Updating asset with data ", transferEvents)

	nfts := []string{address + ":" + event.TokenId.String()}

	fmt.Println("Updating asset metadata for ", nfts)

	updateNFTPrice(nfts)
}

func ParseContractData(address string) {
	assets := GetNFTAssets([]string{address})

	// Gets a list of "contractAddress:tokenId" identifiers for a given contractr
	// Sorted by LOW_PRICE_FIRST
	nftIds := mapToStringArray(assets)

	// Gets asset metadata based on "contractAddress:tokenId" identifiers
	assetsMetadata := GetAssetsMetadata(nftIds)

	// Persist asset metadata in db for each asset
	for _, assetMetadata := range assetsMetadata {
		go UpdateAssetMetadata(assetMetadata)
	}
}

func updateNFTPrice(nfts []string) {

	nftsMetadata := GetAssetsMetadata(nfts)

	for _, meta := range nftsMetadata {
		go UpdateAssetMetadata(meta)
	}
}

func mapToStringArray(assets []models.Asset) []string {
	assetIds := make([]string, len(assets))

	for i := 0; i < len(assets); i++ {
		assetIds[i] = assets[i].Id
	}

	return assetIds
}
