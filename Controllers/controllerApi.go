package Controllers

import (
	"mobile-banking-service/Controllers/User"
	"mobile-banking-service/Dto"
)

type Controller struct {
	Users User.UserControllerInterface
}

func InitControllerApi(u Dto.Utilities) Controller {
	return Controller{
		Users: User.NewController(u),
	}
}
