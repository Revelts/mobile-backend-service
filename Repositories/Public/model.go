package Public

import (
	"mobile-banking-service/Configs"
)

type Repository interface {
	publicCacheRepository
}

type public struct {
	dbCon Configs.DbConInterface
}

func RepositoryNew(dbCon Configs.DbConInterface) Repository {
	return &public{
		dbCon: dbCon,
	}
}
