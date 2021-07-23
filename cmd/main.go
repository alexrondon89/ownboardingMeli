package main

import (
	"ownboardingMeli/internal/api/coingecko_service"
	"ownboardingMeli/internal/server"
	crypto3 "ownboardingMeli/internal/server/controller/crypto"
)

func main () {
	BaseUrl := "https://api.coingecko.com/api/v3/coins/"
	service := coingecko_service.NewCoinGeckoService(BaseUrl)
	controller := crypto3.NewCryptoController(service)
	router := server.GetMeliServer(controller)
	router.Run(":8080")
}
