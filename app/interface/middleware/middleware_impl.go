package middleware

import (
	uc "mygram/app/usecase"
	"mygram/app/usecase/request"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (m *middleware) AuthJwt(ctx *gin.Context) {

	auth := ctx.Request.Header.Get("Authorization")

	if auth == "" {
		resp := map[string]interface{}{
			"message": "Need Sign In",
		}
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	bearer := strings.Split(auth, "Bearer ")

	tokStr := bearer[1]

	tok, err := uc.ValidateToken(tokStr)

	if err != nil {
		resp := map[string]interface{}{
			"message": "Need Sign In",
		}
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	ctx.Set("jwt", tok)
	ctx.Next()
}

func (m *middleware) RegisterUser(c *gin.Context) {
	var req request.RegisterUserReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	c.Set("body", req)
	c.Next()
}

func (m *middleware) LoginUser(c *gin.Context) {
	var req request.LoginUserReq

	err := c.ShouldBind(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	c.Set("body", req)
	c.Next()
}
