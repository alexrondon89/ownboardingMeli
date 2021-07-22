package controller

import "github.com/gin-gonic/gin"

type HomeController interface {
	Initial(c *gin.Context)
}