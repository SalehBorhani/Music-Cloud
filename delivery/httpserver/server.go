package httpserver

import (
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/yazdanbhd/Music-Cloud/config"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver/middleware"
	"github.com/yazdanbhd/Music-Cloud/service/authservice"
	"github.com/yazdanbhd/Music-Cloud/service/totpservice"
)

type Server struct {
	totpSvc totpservice.Service
	authSvc authservice.Service
	cfg     config.Config
}

func New(cfg config.Config, authSvc authservice.Service, totpSvc totpservice.Service) Server {
	return Server{cfg: cfg, authSvc: authSvc, totpSvc: totpSvc}
}

func (s Server) Run() {
	e := echo.New()
	// Use echo middlewares
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	userGroup := e.Group("/api/users")
	userGroup.POST("/register", s.UserRegister)
	userGroup.POST("/login", s.UserLogin)

	musicGroup := e.Group("/api/music")
	musicGroup.POST("/upload", s.UploadMusic, middleware.Auth(s.authSvc, s.cfg.AuthConfig))

	e.Logger.Fatal(e.Start(":8080"))
}
