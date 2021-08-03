package coingecko_controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ownboardingMeli/internal/server/controller/coingecko_controller/dto"
	"ownboardingMeli/internal/service"
	"ownboardingMeli/pkg/Errors"
)

type CryptoController struct {
	CryptoService service.CryptoService
}

func NewCryptoController(service service.CryptoService) *CryptoController {
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

	if err := c.BindQuery(&input); err !=nil{
		log.Println("ERROR OCCURRED: ", err)
		c.JSON(http.StatusBadRequest, Errors.BuildBadRequestError(err.Error()))
		return
	}

	list, _ := cr.CryptoService.GetListPrice(input.Coins, input.Currency)

	c.JSON(http.StatusOK, list)
}


