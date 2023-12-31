package blacklist

import (
	"context"

	"github.com/MaxFando/anti-bruteforce/internal/domain/network"
)

type Service interface {
	AddIP(ctx context.Context, network network.IPNetwork) error
	RemoveIP(ctx context.Context, network network.IPNetwork) error
	GetIPList(ctx context.Context) ([]network.IPNetwork, error)
}

type UseCase struct {
	blackListService Service
}

func NewUseCase(blackListService Service) *UseCase {
	return &UseCase{blackListService: blackListService}
}

func (u *UseCase) AddIP(ctx context.Context, network network.IPNetwork) error {
	return u.blackListService.AddIP(ctx, network)
}

func (u *UseCase) RemoveIP(ctx context.Context, network network.IPNetwork) error {
	return u.blackListService.RemoveIP(ctx, network)
}

func (u *UseCase) GetIPList(ctx context.Context) ([]network.IPNetwork, error) {
	return u.blackListService.GetIPList(ctx)
}
