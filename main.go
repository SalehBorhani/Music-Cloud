package main

import (
	"github.com/yazdanbhd/Music-Cloud/config"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver"
	"github.com/yazdanbhd/Music-Cloud/service/authservice"
)

func main() {

	// Setup configuration
	cfg := config.New("config.yml")

	authSvc := authservice.New(cfg.AuthConfig)
	server := httpserver.New(cfg, authSvc)

	server.Run()
}
