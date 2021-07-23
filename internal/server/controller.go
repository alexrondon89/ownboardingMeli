package server

import "github.com/gin-gonic/gin"

type CryptoController interface {
	Request(c *gin.Context)
}