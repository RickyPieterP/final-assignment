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
}
