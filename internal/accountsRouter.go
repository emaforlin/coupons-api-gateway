package internal

import (
	"context"
	"log"

	accountsv1 "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"google.golang.org/grpc"
)

type accountRouterImpl struct {
	client accountsv1.AccountsClient
}

func (a *accountRouterImpl) SignupPerson(ctx context.Context, in *accountsv1.AddPersonAccountRequest) error {
	_, err := a.client.AddPersonAccount(ctx, in)
	return err
}

func NewAccountRouter(host string, grpcOpts ...grpc.DialOption) AccountRouter {
	conn, err := grpc.NewClient(host, grpcOpts...)
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	return &accountRouterImpl{
		client: accountsv1.NewAccountsClient(conn),
	}
}
