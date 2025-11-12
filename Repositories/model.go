package Repositories

import (
	"mobile-banking-service/Configs"
	"mobile-banking-service/Repositories/User"
)

type Repository struct {
	User User.Repository
}

func InitRepo(dbCon Configs.DbConInterface) Repository {
	return Repository{
		User: User.RepositoryNew(dbCon),
	}
}
