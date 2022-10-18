package middleware

import (
	"fmt"
	"mygram/app/usecase/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (m *middleware) RegisterUser(c *gin.Context) {
	var req request.RegisterUserReq

	err := c.ShouldBindJSON(&req)
	fmt.Println(err)
	if err != nil {
		c.Abort()
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		c.Abort()
	}

	c.Set("body", req)
	c.Next()
}

func (m *middleware) LoginUser(c *gin.Context) {
	var req request.LoginUserReq

	err := c.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		c.Abort()
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		c.Abort()
		return
	}

	c.Set("body", req)
	c.Next()
}
