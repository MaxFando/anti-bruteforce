package grpc

import (
	"context"
	"github.com/MaxFando/anti-bruteforce/internal/delivery/grpc/authpb"
	"github.com/MaxFando/anti-bruteforce/internal/domain/network"
	"github.com/MaxFando/anti-bruteforce/internal/usecase/auth"
	"github.com/MaxFando/anti-bruteforce/pkg/utils"
)

type AuthServer struct {
	authpb.AuthorizationServer
	uc *auth.UseCase
}

func NewAuthServer(uc *auth.UseCase) *AuthServer {
	return &AuthServer{uc: uc}
}

func (s *AuthServer) TryAuthorization(ctx context.Context, req *authpb.AuthorizationRequest) (*authpb.AuthorizationResponse, error) {
	utils.Logger.Info("Try Authorization by GRPC")

	request, err := network.NewRequest(
		req.GetRequest().GetLogin(),
		req.GetRequest().GetPassword(),
		req.GetRequest().GetIp(),
	)

	if err != nil {
		return nil, err
	}

	isAllowed, err := s.uc.TryAuthorization(ctx, request)
	if err != nil {
		return nil, err
	}

	return &authpb.AuthorizationResponse{Ok: isAllowed}, nil
}
