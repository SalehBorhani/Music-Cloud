package main

import (
	"github.com/yazdanbhd/Music-Cloud/config"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver"
	"github.com/yazdanbhd/Music-Cloud/service/authservice"
	"github.com/yazdanbhd/Music-Cloud/service/totpservice"
)

func main() {

	// Setup configuration
	cfg := config.New("config.yml")

	authSvc := authservice.New(cfg.AuthConfig)
	totpConfig := totpservice.Config{
		AppName: "Music-Cloud",
	}
	totpSvc := totpservice.New(totpConfig)
	server := httpserver.New(cfg, authSvc, totpSvc)

	server.Run()
}
