package crypto

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ownboardingMeli/internal/api"
	"ownboardingMeli/internal/server/controller/crypto/dto"
	"ownboardingMeli/pkg/Errors"
)

type CryptoController struct {
	CryptoService api.CryptoService
}

func NewCryptoController(service api.CryptoService) *CryptoController {
	return &CryptoController{CryptoService: service}
}

func (cr *CryptoController) CoinPrice(c *gin.Context){
	var input dto.InputCoin

	if err := c.BindQuery(&input); err !=nil{
		log.Println("ERROR OCCURRED: ", err)
		c.JSON(http.StatusBadRequest, Errors.BuildBadRequestError(err.Error()))
		return
	}

	response, err := cr.CryptoService.GetPrice(input.Coin, input.Currency)
	if err != nil {
		log.Println("ERROR OCCURRED GETTING COIN PRICE: ", err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if response.Content == nil{
		log.Println("CONTENT IS NOT PRESENT IN RESPONSE")

		c.JSON(http.StatusPartialContent, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (cr *CryptoController) ListPrice(c *gin.Context){
	var input dto.InputListCoin
	 ListCoins := []string{"bitcoin","ethereum", "cardano"}

	if err := c.ShouldBindQuery(&input); err !=nil{
		log.Println("ERROR OCCURRED: ", err)
		c.JSON(http.StatusBadRequest, Errors.BuildBadRequestError(err.Error()))
		return
	}

	list, _ := cr.CryptoService.GetListPrice(ListCoins, input.Currency)

	c.JSON(http.StatusOK, list)
}


