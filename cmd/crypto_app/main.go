package main

import (
	cli "ownboardingMeli/internal/client/coingecko/http"
	contr "ownboardingMeli/internal/server/controller/coingecko"
	"ownboardingMeli/internal/server/router"
	serv "ownboardingMeli/internal/service/coingecko"
)

func main () {
	BaseUrl := "https://api.coingecko.com/api/v3/coins/"
	coinGeckoClient := cli.NewCoinGeckoClient(BaseUrl)
	service := serv.NewCoinGeckoService(coinGeckoClient)
	controller := contr.NewCryptoController(service)
	router := router.GetRouter(controller)
	router.Run(":8080")
}
