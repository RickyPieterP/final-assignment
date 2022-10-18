package server

import (
	"fmt"
	"mygram/app/interface/container"
	"mygram/app/interface/middleware"
	"mygram/app/shared/config"
	"mygram/app/transport"

	"github.com/gin-gonic/gin"
)

func StartServer(container container.Container) {
	app := gin.New()
	app.Use(gin.Recovery())

	// setup middleware
	middleware := middleware.NewMiddleware(&container)
	// setup transport
	transport := transport.SetupTransport(&container)
	setupRouter(transport, middleware, app)

	// run gin apps
	fmt.Println(app.Run(config.Server.PORTHTTP))
}
