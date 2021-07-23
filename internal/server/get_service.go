package server

import (
	"github.com/gin-gonic/gin"
)

func GetMeliServer(input CryptoController) *gin.Engine{
	router := gin.Default()
	m := router.Group("/meli")
	m.GET("/myapi", input.Request)
	return router
}
