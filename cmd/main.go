package main

import (
	"ownboardingMeli/internal/api/coingecko_service"
	"ownboardingMeli/internal/client/coingecko_client"
	"ownboardingMeli/internal/server"
	crypto3 "ownboardingMeli/internal/server/controller/crypto"
)

func main () {
	BaseUrl := "https://api.coingecko.com/api/v3/coins/"
	coinGeckoClient := coingecko_client.NewCoinGeckoClient(BaseUrl)
	service := coingecko_service.NewCoinGeckoService(coinGeckoClient)
	controller := crypto3.NewCryptoController(service)
	router := server.GetMeliServer(controller)
	router.Run(":8080")
}
