package server

import (
	"github.com/gin-gonic/gin"
)

func GetMeliServer(input CryptoController) *gin.Engine{
	router := gin.Default()
	m := router.Group("/meli")
	m.GET("/coinprice", input.CoinPrice)
	m.GET("/listprice", input.ListPrice)
	return router
}
