package coingecko

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ownboardingMeli/internal/service"
	errservice "ownboardingMeli/pkg/errors/service"
)

type CryptoController struct {
	CryptoService service.CryptoService
}

type InputCoin struct {
	Coin 			string		`json:"coin" validate:"required"`
	Money						`json:"money" validate:"required"`
}

type InputListCoin struct {
	Coins			[]string	`json:"coins" validate:"required"`
	Money						`json:"money" validate:"required"`
}

type Money struct {
	Currency string
}

func NewCryptoController(service service.CryptoService) *CryptoController {
	return &CryptoController{CryptoService: service}
}

func (cr *CryptoController) CoinPrice(c *gin.Context){
	var input InputCoin

	if err := c.BindQuery(&input); err !=nil{
		log.Println("ERROR OCCURRED: ", err)
		c.JSON(http.StatusBadRequest, errservice.BuildBadRequestError(err.Error()))
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
	var input InputListCoin

	if err := c.BindQuery(&input); err !=nil{
		log.Println("ERROR OCCURRED: ", err)
		c.JSON(http.StatusBadRequest, errservice.BuildBadRequestError(err.Error()))
		return
	}

	list, _ := cr.CryptoService.GetListPrice(input.Coins, input.Currency)

	c.JSON(http.StatusOK, list)
}


