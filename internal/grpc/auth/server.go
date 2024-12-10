package auth

import (
	"context"

	ssov1 "github.com/MirzaDgtu/firstcode_protos"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Login(ctx context.Context,
		email string,
		password string,
		appID int) (token string, err error)

	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
		first_name string,
		name string,
		last_name string,
		phone string,
		sex string) (userID int64, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	// TODO
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	// TODO
}
