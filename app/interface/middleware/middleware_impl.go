package middleware

import (
	"fmt"
	"mygram/app/usecase/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
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

func (m *middleware) ValidateUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	tokenMap := map[string]string{
		"user_id": "",
		"user": "",
	}
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return "", nil	
	})
	fmt.Println(tokenString)
	if err != nil {
		fmt.Println(err)
	}

	if claims, err := tokenString.Claims.(jwt.MapClaims); err && tokenString.Valid {
		fmt.Println("claims", claims)
		for key, val := range claims {
			if s, ok := val.(string); ok {
				tokenMap[key] = s
			}
		}
	}

	c.Set("user_id", tokenMap["user_id"])
}