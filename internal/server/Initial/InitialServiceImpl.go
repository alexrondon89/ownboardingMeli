package Initial

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ownboardingMeli/internal/models"
)

type Init struct {
}

func NewInit() *Init{
	return &Init{}
}

func (i *Init) Initial (c *gin.Context){
	print("llego hasta el initial impl")

	var input models.Input
	var output models.Response
	if err := c.BindQuery(&input); err !=nil{
		output.Message = err.Error()

		c.JSON(http.StatusBadRequest, output)
		return
	}

	output.Message =  input.Data
	c.JSON(http.StatusOK, output)
}


