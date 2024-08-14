package config

import (
	"fmt"
	"os"

	"cosmossdk.io/errors"
	"github.com/lpernett/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Port       = "8080"
	BaseURL    = os.Getenv("BASE_URL")
	ListenAddr = os.Getenv("LISTEN_ADDR")

	AccountsBaseUrl = "/accounts"
)

func LoadConfig() {
	godotenv.Load()
	if os.Getenv("PORT") != "" {
		Port = os.Getenv("PORT")
	}
}

func MustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func MustConnGRPC(conn **grpc.ClientConn, addr string) {
	var err error

	// *conn, err = grpc.DialContext(ctx, addr)
	*conn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
