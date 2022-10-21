package middleware

import (
	"mygram/app/interface/container"

	"github.com/gin-gonic/gin"
)

type middleware struct {
}

type Middleware interface {
	//jwt
	AuthJwt(c *gin.Context)
	// User
	RegisterUser(*gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	// Photo
	CreatePhoto(c *gin.Context)
	FindPhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)
	// Social Media
	CreateSocialMedia(c *gin.Context)
	FindSocialMedia(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	DeleteSocialMedia(c *gin.Context)
}

func NewMiddleware(container *container.Container) Middleware {
	return &middleware{}
}
