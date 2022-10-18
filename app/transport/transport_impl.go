package transport

import (
	"mygram/app/usecase"
	"mygram/app/usecase/request"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	usecase  usecase.Usecase
	response string
}

func NewTransport(usecase usecase.Usecase, response string) *Transport {
	return &Transport{
		usecase:  usecase,
		response: response,
	}
}

func (t *Transport) RegisterUser(c *gin.Context) {
	body := c.MustGet("body").(request.RegisterUserReq)

	res, httpStatus := t.usecase.RegisterUser(body)

	c.JSON(httpStatus, res)
}
