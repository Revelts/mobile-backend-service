package User

import (
	"mobile-banking-service/Dto"
	"mobile-banking-service/Library/Helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	Ping(g *gin.Context)
}

func (u user) Ping(g *gin.Context) {
	Helper.HttpResponseSuccess(g, http.StatusOK)
}

func NewController(u Dto.Utilities) UserControllerInterface {
	return &user{u}
}
