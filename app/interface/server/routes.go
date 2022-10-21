package server

import (
	"mygram/app/interface/middleware"
	"mygram/app/transport"

	"github.com/gin-gonic/gin"
)

func setupRouter(transport *transport.Tp, middleware middleware.Middleware, app *gin.Engine) {
	ping := app.Group("/ping")
	ping.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	users := app.Group("/users")
	users.POST("/register", middleware.RegisterUser, transport.Transport.RegisterUser)
	users.POST("/login", middleware.LoginUser, transport.Transport.LoginUser)
	users.PUT("/update", middleware.AuthJwt, middleware.UpdateUser, transport.Transport.UpdateUser)
	users.DELETE("/delete", middleware.AuthJwt, middleware.DeleteUser, transport.Transport.DeleteUser)
	photo := app.Group("/photo")
	photo.POST("/create", middleware.AuthJwt, middleware.CreatePhoto, transport.Transport.CreatePhoto)
	photo.GET("/", middleware.AuthJwt, middleware.FindPhoto, transport.Transport.FindPhoto)
	photo.PUT("/:photoId", middleware.AuthJwt, middleware.UpdatePhoto, transport.Transport.UpdatePhoto)
	photo.DELETE("/:photoId", middleware.AuthJwt, middleware.DeletePhoto, transport.Transport.DeletePhoto)
	socmed := app.Group("/socialmedias")
	socmed.POST("/POST", middleware.AuthJwt, middleware.CreateSocialMedia, transport.Transport.CreateSocialMedia)
	// socmed.GET("/", middleware.AuthJwt, middleware.FindSocialMedia, transport.Transport.SocialMedia)
}
