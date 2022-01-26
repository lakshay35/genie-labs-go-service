package main

import (
	"fmt"

	"github.com/genie-service/services"
	"github.com/genie-service/utils/database"
	"github.com/genie-service/utils/environment"
	"github.com/gin-gonic/gin"
)

func main() {
	addresses := []string{
		"0x60e4d786628fea6478f785a6d7e704777c86a7c6", // MAYC
		"0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d", // BAYC
	}

	database.InitializeDatabase()

	for _, address := range addresses {
		fmt.Println("Hydrating database with intial 50 NFT Pricing data sorted by LOW_PRICE_FIRST\n")
		go services.ParseContractData(address)

		fmt.Println("Initiating WSS over RPC to an Ethereum Node for contract ", address, "\n")
		go services.SubscribeToContractEvents(address)
	}

	r := gin.Default()

	err := r.Run(":" + environment.GetEnvVariable("PORT"))
	if err != nil {
		panic("unable to start server")
	}
}
