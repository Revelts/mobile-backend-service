package Routes

import (
	"mobile-banking-service/Controllers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Controller Controllers.Controller
	Gin        *gin.Engine
}

func (app *Routes) CollectRoutes() *gin.Engine {
	appRoute := app.Gin
	appRoute.GET("/ping", app.Controller.Users.Ping)
	return appRoute
}
