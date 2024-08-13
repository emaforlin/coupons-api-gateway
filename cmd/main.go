package main

import (
	"os"

	"github.com/emaforlin/api-gateway/internal/config"
	"github.com/emaforlin/api-gateway/internal/entities"
	"github.com/emaforlin/api-gateway/internal/middlewares"
	"github.com/emaforlin/api-gateway/internal/server"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lpernett/godotenv"
)

const (
	port = "8080"

	accountsBaseUrl = "/accounts"
)

func main() {
	godotenv.Load()

	svc := new(server.APIGatewayServer)

	baseUrl := os.Getenv("BASE_URL")
	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}
	addr := os.Getenv("LISTEN_ADDR")

	config.MustMapEnv(&svc.AccountSvcAddr, "ACCOUNTS_SERVICE_ADDR")

	config.MustConnGRPC(&svc.AccountSvcConn, svc.AccountSvcAddr)

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger())

	// main router
	router := e.Group(baseUrl)

	// private router
	priv := router.Group("/_")
	priv.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		Skipper: func(c echo.Context) bool {
			return c.Path() == baseUrl+"/login"
		},
	}))

	router.POST(accountsBaseUrl+"/signup", svc.SignupHandler)
	router.POST(accountsBaseUrl+"/signup/partner", svc.SignupPartnerHandler)
	router.POST("/login", svc.LoginHandler)

	priv.GET("/priv", func(c echo.Context) error {
		return c.JSON(200, "hola")
	}, middlewares.AllowedRoles("FoodPlace", "Customer"))

	e.Logger.Fatal(e.Start(addr + ":" + srvPort))
}
