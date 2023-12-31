package providers

import (
	"github.com/MaxFando/anti-bruteforce/internal/repository/inmemory/bucket"
	"github.com/MaxFando/anti-bruteforce/internal/repository/postgres/list"
	"github.com/MaxFando/anti-bruteforce/internal/store/inmemory"
	"github.com/MaxFando/anti-bruteforce/internal/store/postgres"
)

type RepositoryProvider struct {
	db *postgres.Client

	ipBucketRepo       *bucket.Repository
	loginBucketRepo    *bucket.Repository
	passwordBucketRepo *bucket.Repository

	blackListRepo *list.BlackListRepository
	whiteListRepo *list.WhiteListRepository
}

func NewRepositoryProvider(db *postgres.Client) *RepositoryProvider {
	return &RepositoryProvider{
		db: db,
	}
}

func (rp *RepositoryProvider) BootPrefixProviderContract() {
	rp.ipBucketRepo = bucket.NewRepository(inmemory.New())
	rp.loginBucketRepo = bucket.NewRepository(inmemory.New())
	rp.passwordBucketRepo = bucket.NewRepository(inmemory.New())

	rp.blackListRepo = list.NewBlackListRepository(rp.db)
	rp.whiteListRepo = list.NewWhiteListRepository(rp.db)
}
