package main

import (
	"ownboardingMeli/internal/client/coingecko_client"
	"ownboardingMeli/internal/server/controller/coingecko_controller"
	"ownboardingMeli/internal/server/router"
	"ownboardingMeli/internal/service/coingecko_service"
)

func main () {
	BaseUrl := "https://api.coingecko.com/api/v3/coins/"
	coinGeckoClient := coingecko_client.NewCoinGeckoClient(BaseUrl)
	service := coingecko_service.NewCoinGeckoService(coinGeckoClient)
	controller := coingecko_controller.NewCryptoController(service)
	router := router.GetRouter(controller)
	router.Run(":8080")
}
