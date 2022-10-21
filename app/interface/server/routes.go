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
	photo := app.Group("/photo")
	photo.POST("/create", middleware.AuthJwt, middleware.CreatePhoto, transport.Transport.CreatePhoto)
	photo.GET("/", middleware.AuthJwt, middleware.FindPhoto, transport.Transport.FindPhoto)
	photo.PUT("/:photoId", middleware.AuthJwt, middleware.UpdatePhoto, transport.Transport.UpdatePhoto)
	photo.DELETE("/:photoId", middleware.AuthJwt, middleware.DeletePhoto, transport.Transport.DeletePhoto)

	comment := app.Group("/comment")
	comment.POST("/", middleware.AuthJwt, middleware.CreateComment, transport.Transport.CreateComment)
	comment.GET("/", middleware.AuthJwt, middleware.FindComment, transport.Transport.FindComment)
	comment.PUT("/:commentId", middleware.AuthJwt, middleware.UpdateComment, transport.Transport.UpdateComment)
	comment.DELETE("/:commentId", middleware.AuthJwt, middleware.DeleteComment, transport.Transport.DeleteComment)

	socmed := app.Group("/socialmedias")
	socmed.POST("/POST", middleware.AuthJwt, middleware.CreateSocialMedia, transport.Transport.CreateSocialMedia)
	// socmed.GET("/", middleware.AuthJwt, middleware.FindSocialMedia, transport.Transport.SocialMedia)
}
