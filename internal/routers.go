package internal

import (
	"context"

	accountPb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
)

type AccountRouter interface {
	SignupPerson(ctx context.Context, in *accountPb.AddPersonAccountRequest) error
}
