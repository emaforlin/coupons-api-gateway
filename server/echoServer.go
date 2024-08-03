package server

import (
	"github.com/emaforlin/api-gateway/config"
	"github.com/emaforlin/api-gateway/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type echoServer struct {
	cfg *config.Config
	app *echo.Echo
}

func (s *echoServer) Start() {
	s.initializeHandlers()
	s.app.Use(middleware.Recover(), middleware.Logger())

	addr := ":80"
	if s.cfg.App.Debug {
		addr = ":8080"
	}

	s.app.Logger.Fatal(s.app.Start(addr))
}

func NewHttpServer(conf *config.Config) Server {
	return &echoServer{
		cfg: conf,
		app: echo.New(),
	}
}

func (s *echoServer) initializeHandlers() {
	accountsRouter := internal.NewAccountRouter(s.cfg.Services["accounts"].Host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	accountsHandler := internal.NewAccountHandler(accountsRouter)

	public := s.app.Group("/" + s.cfg.App.Version)

	public.POST("/signup", accountsHandler.SignupPerson)
}
