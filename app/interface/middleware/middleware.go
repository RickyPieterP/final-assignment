package middleware

import (
	"mygram/app/interface/container"

	"github.com/gin-gonic/gin"
)

type middleware struct {
}

type Middleware interface {
	RegisterUser(*gin.Context)
	LoginUser(c *gin.Context)
	AuthJwt(c *gin.Context)
}

func NewMiddleware(container *container.Container) Middleware {
	return &middleware{}
}
