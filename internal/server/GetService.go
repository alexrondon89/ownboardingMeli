package server

import (
	"github.com/gin-gonic/gin"
	"ownboardingMeli/internal/server/controller"
)

func GetMeliServer(input controller.HomeController) *gin.Engine{
	router := gin.Default()
	m := router.Group("/meli")
	m.GET("/myapi", input.Initial)
	return router
}
