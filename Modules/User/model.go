package User

import (
	ModulesDto "mobile-banking-service/Modules/Dto"
	"mobile-banking-service/Repositories"
)

type user struct {
	repo Repositories.Repository
}

type UserModules interface {
}

func NewModules(moduleConfig ModulesDto.ModulesConfig) UserModules {
	return &user{
		repo: moduleConfig.Repo,
	}
}
