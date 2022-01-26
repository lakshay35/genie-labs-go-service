package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/genie-service/models"
	"github.com/genie-service/utils/database"
)

func GetNFTAssets(collections []string) []models.Asset {
	payload := models.GetNFTAssetsPayload{
		Filter: models.Filter{
			VerifiedOnly: false,
			Sort:         "LOW_PRICE_FIRST",
			Collections:  collections,
			Traits:       []string{},
			NSFW:         true,
		},
		Size: 50,
	}

	requestPayload := marshalPayload(payload)

	httpResp, httpErr := http.Post(
		"https://api-mainnet.rarible.com/marketplace/search/v1/items",
		"application/json",
		bytes.NewBuffer(requestPayload),
	)

	if httpErr != nil {
		panic(httpErr)
	}

	defer httpResp.Body.Close()

	body, respErr := ioutil.ReadAll(httpResp.Body)
	if respErr != nil {
		panic(respErr)
	}

	var assets []models.Asset
	unmarshalErr := json.Unmarshal(body, &assets)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	return assets
}

func GetAssetsMetadata(nfts []string) []models.NFTAssetMetadata {

	requestPayload := marshalPayload(nfts)

	httpResp, httpErr := http.Post(
		"https://api-mainnet.rarible.com/marketplace/api/v4/items/map",
		"application/json",
		bytes.NewBuffer(requestPayload),
	)

	if httpErr != nil {
		panic(httpErr)
	}

	defer httpResp.Body.Close()

	body, respErr := ioutil.ReadAll(httpResp.Body)

	if respErr != nil {
		panic(respErr)
	}

	var assetsMetadata []models.NFTAssetMetadata

	unmarshalErr := json.Unmarshal(body, &assetsMetadata)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	return assetsMetadata
}

func UpdateAssetMetadata(metadata models.NFTAssetMetadata) {
	connection := database.GetConnection()
	defer database.CloseConnection(connection)
	query := `INSERT INTO assets (id, token_id, collection_address, collection_name) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING;`

	stmt := database.PrepareStatement(connection, query)

	collectionName := strings.Split(metadata.Properties.Name, " ")[0]

	identifiers := strings.Split(metadata.Id, ":")

	_, queryErr := stmt.Exec(
		metadata.Id,
		identifiers[1],
		identifiers[0],
		collectionName,
	)

	if queryErr != nil {
		panic(queryErr)
	}

	query = `INSERT INTO asset_price (id, status, current_offer_id, curr_price_usd, curr_price_jpy, curr_price_eth) VALUES (
		$1, $2, $3, $4, $5, $6) ON CONFLICT(id) DO UPDATE SET status = $2, current_offer_id = $3, curr_price_usd = $4, curr_price_jpy = $5, curr_price_eth = $6`

	stmt = database.PrepareStatement(connection, query)

	offerId := ""
	currPrice := 0.0
	status := "NOT_FOR_SALE"

	if metadata.Item.Ownership != nil {
		status = metadata.Item.Ownership.Status
		currPrice = metadata.Item.Ownership.PriceEth
	}

	if metadata.Item.Offer != nil {
		offerId = metadata.Item.Offer.Id
		status = "HAS_OFFER"
		currPrice = metadata.Item.Offer.BuyPrice
	}

	multiCurrencyPrices := getMultiCurrencyPriceOfEth(currPrice)

	_, execErr := stmt.Exec(
		metadata.Id,
		status,
		offerId,
		multiCurrencyPrices.USDPrice,
		multiCurrencyPrices.JPYPrice,
		multiCurrencyPrices.ETHPrice,
	)

	if execErr != nil {
		panic(execErr)
	}
}

func marshalPayload(payload interface{}) []byte {
	json, marshalErr := json.Marshal(payload)

	if marshalErr != nil {
		panic(marshalErr)
	}

	return json
}

func getMultiCurrencyPriceOfEth(price float64) models.MultiCurrencyPrices {
	quote := getEthExchangeRates()

	usdQuote, usdQuoteErr := strconv.ParseFloat(quote.Data.Rates["USD"], 64)
	jpyQuote, jpyQuoteErr := strconv.ParseFloat(quote.Data.Rates["JPY"], 64)

	if usdQuoteErr != nil {
		panic(usdQuoteErr)
	}

	if jpyQuoteErr != nil {
		panic(jpyQuoteErr)
	}

	return models.MultiCurrencyPrices{
		USDPrice: price * usdQuote,
		JPYPrice: price * jpyQuote,
		ETHPrice: price,
	}
}

func getEthExchangeRates() models.EthereumExchangeQuote {
	httpResp, httpErr := http.Get(
		"https://api.coinbase.com/v2/exchange-rates?currency=ETH",
	)

	if httpErr != nil {
		panic(httpErr)
	}

	defer httpResp.Body.Close()

	body, bodyErr := ioutil.ReadAll(httpResp.Body)

	if bodyErr != nil {
		panic(bodyErr)
	}

	var quote models.EthereumExchangeQuote

	unmarshalErr := json.Unmarshal(body, &quote)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	return quote
}

//  https://api.coinbase.com/v2/exchange-rates?currency=ETH - API To get eth conversion rate to USD and CAD
