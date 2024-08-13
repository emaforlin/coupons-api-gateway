package server

import (
	"google.golang.org/grpc"
)

type APIGatewayServer struct {
	AccountSvcAddr string
	AccountSvcConn *grpc.ClientConn
}
