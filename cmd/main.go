package main

import (
	"os"

	"github.com/emaforlin/api-gateway/internal/config"
	"github.com/emaforlin/api-gateway/internal/entities"
	"github.com/emaforlin/api-gateway/internal/middlewares"
	"github.com/emaforlin/api-gateway/internal/server"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()
	svc := new(server.APIGatewayServer)

	config.MustMapEnv(&svc.AccountSvcAddr, "ACCOUNTS_SERVICE_ADDR")

	config.MustConnGRPC(&svc.AccountSvcConn, svc.AccountSvcAddr)

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger())

	// main router
	router := e.Group(config.BaseURL)
	router.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: middlewares.JwtSkipperFunc,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	router.POST(config.AccountsBaseUrl+"/signup", svc.SignupHandler)
	router.POST(config.AccountsBaseUrl+"/signup/partner", svc.SignupPartnerHandler)
	router.POST("/login", svc.LoginHandler)

	router.GET("/priv", func(c echo.Context) error {
		return c.JSON(200, "hola")
	}, middlewares.AllowedRoles("FoodPlace", "Customer"))

	e.Logger.Fatal(e.Start(config.ListenAddr + ":" + config.Port))
}
