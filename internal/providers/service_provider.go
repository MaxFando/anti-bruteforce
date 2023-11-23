package providers

import (
	"github.com/MaxFando/anti-bruteforce/internal/service/blacklist"
	"github.com/MaxFando/anti-bruteforce/internal/service/bucket"
	"github.com/MaxFando/anti-bruteforce/internal/service/whitelist"
)

type ServiceProvider struct {
	BlacklistService *blacklist.Service
	WhitelistService *whitelist.Service
	BucketService    *bucket.Service
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) RegisterDependencies(repoProvider *RepositoryProvider) {
	sp.BlacklistService = blacklist.NewService(repoProvider.blackListRepo)
	sp.WhitelistService = whitelist.NewService(repoProvider.whiteListRepo)

	sp.BucketService = bucket.NewService(
		repoProvider.ipBucketRepo,
		repoProvider.loginBucketRepo,
		repoProvider.passwordBucketRepo,
	)
}