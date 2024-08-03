package main

import (
	config "github.com/emaforlin/api-gateway/config"
	srv "github.com/emaforlin/api-gateway/server"
)

func main() {
	config.InitViper("config")
	cfg := config.LoadConfig()

	server := srv.NewHttpServer(cfg)
	server.Start()
}
