package router

import (
	"github.com/gin-gonic/gin"
	"ownboardingMeli/internal/server"
)

func GetRouter(input server.CryptoController) *gin.Engine{
	router := gin.Default()
	m := router.Group("/meli")
	m.GET("/coinprice", input.CoinPrice)
	m.GET("/listprice", input.ListPrice)
	return router
}
