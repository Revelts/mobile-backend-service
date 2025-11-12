package User

import "mobile-banking-service/Configs"

type Repository interface {
}

// CONSTRUCTOR STRUCT
type user struct {
	dbCon Configs.DbConInterface
}

// CONSTRUCTOR FUNCTION FOR USER REPOSITORY
func RepositoryNew(dbCon Configs.DbConInterface) Repository {
	return &user{
		dbCon: dbCon,
	}
}
