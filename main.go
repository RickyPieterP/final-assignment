package main

import (
	"mygram/app/interface/container"
	"mygram/app/interface/server"
)

func main() {
	server.StartServer(container.SetupContainer())
}
