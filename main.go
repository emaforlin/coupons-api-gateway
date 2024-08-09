package main

import (
	"fmt"
	"os"

	"cosmossdk.io/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lpernett/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "8080"
)

type apigatewayServer struct {
	accountSvcAddr string
	accountSvcConn *grpc.ClientConn
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(errors.Wrapf(err, "could not load .env"))
	}

	svc := new(apigatewayServer)

	baseUrl := os.Getenv("BASE_URL")

	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}

	addr := os.Getenv("LISTEN_ADDR")

	mustMapEnv(&svc.accountSvcAddr, "ACCOUNT_SERVICE_ADDR")

	mustConnGRPC(&svc.accountSvcConn, svc.accountSvcAddr)

	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger())

	router := e.Group(baseUrl)

	router.GET("/signup", svc.signupHandler)

	e.Logger.Fatal(e.Start(addr + ":" + srvPort))
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(conn **grpc.ClientConn, addr string) {
	var err error

	// *conn, err = grpc.DialContext(ctx, addr)
	*conn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
