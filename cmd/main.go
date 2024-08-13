package main

import (
	"os"

	"cosmossdk.io/errors"
	"github.com/emaforlin/api-gateway/internal/config"
	"github.com/emaforlin/api-gateway/internal/entities"
	"github.com/emaforlin/api-gateway/internal/middlewares"
	"github.com/emaforlin/api-gateway/internal/server"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lpernett/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "8080"

	accountsBaseUrl = "/accounts"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(errors.Wrapf(err, "could not load .env"))
	}

	svc := new(server.APIGatewayServer)

	baseUrl := os.Getenv("BASE_URL")

	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}

	addr := os.Getenv("LISTEN_ADDR")

	config.MustMapEnv(&svc.AccountSvcAddr, "ACCOUNT_SERVICE_ADDR")

	mustConnGRPC(&svc.AccountSvcConn, svc.AccountSvcAddr)

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger())

	// main router
	router := e.Group(baseUrl)

	// private router
	restricted := router.Group("")
	restricted.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	accountRouter := router.Group(accountsBaseUrl)

	accountRouter.POST("/signup", svc.SignupHandler)
	accountRouter.POST("/signup/partner", svc.SignupPartnerHandler)

	restricted.GET("/priv", func(c echo.Context) error {
		return c.JSON(200, "hola")
	}, middlewares.AllowedRoles("FoodPlace", "Customer"))

	e.Logger.Fatal(e.Start(addr + ":" + srvPort))
}

func mustConnGRPC(conn **grpc.ClientConn, addr string) {
	var err error

	// *conn, err = grpc.DialContext(ctx, addr)
	*conn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
