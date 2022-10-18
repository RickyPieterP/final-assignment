package middleware

import (
	"fmt"
	"mygram/app/usecase/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func (m *middleware) RegisterUser(c *gin.Context) {
	var req request.RegisterUserReq

	err := c.ShouldBindJSON(&req)
	fmt.Println(err)
	if err != nil {
		// log.Error(c, err)
		// m.responseWriter.Error(c, http.StatusOK, config.RCMobeDataNotValid)
		c.Abort()
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		// log.Error(c, err)
		// m.responseWriter.Error(c, http.StatusOK, config.RCMobeDataNotValid)
		c.Abort()
	}

	c.Set("body", req)
	c.Next()
}
