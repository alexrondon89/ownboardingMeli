package crypto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ownboardingMeli/internal/api"
	"ownboardingMeli/internal/server/controller/crypto/dto"
)

type CryptoController struct {
	CryptoService api.CryptoService
}

func NewCryptoController(service api.CryptoService) *CryptoController {
	return &CryptoController{CryptoService: service}
}

func (cr *CryptoController) Request(c *gin.Context){
	var data dto.Input

	if err := c.BindQuery(&data); err !=nil{
		c.JSON(http.StatusPartialContent, dto.BuildPartialResponse(data.Id))
		return
	}

	response, err := cr.CryptoService.GetPrice(data.Id, data.Currency)

	if err != nil {
		c.JSON(http.StatusPartialContent, dto.BuildPartialResponse(data.Id))
		return
	}

	c.JSON(http.StatusOK, response)
}


