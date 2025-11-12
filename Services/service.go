package Services

import (
	"flag"
	"mobile-banking-service/Configs"
	"mobile-banking-service/Constants"
	"mobile-banking-service/Controllers"
	"mobile-banking-service/Dto"
	"mobile-banking-service/Modules"
	ModulesDto "mobile-banking-service/Modules/Dto"
	"mobile-banking-service/Modules/User"
	"mobile-banking-service/Repositories"
	"mobile-banking-service/Routes"

	"github.com/gin-gonic/gin"
)

var AppEnv = flag.String("env", "", "define environment")
var routesConfig *gin.Engine

func init() {
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	routesConfig = gin.Default()
	if *AppEnv == "" {
		*AppEnv = Constants.Localhost
	}
}

// AppInitialization Application Engine initialization
func AppInitialization() {
	newConfig := Configs.GetEnvironment(*AppEnv).LoadConfig()
	connection := newConfig.Database.BuildConnection()
	ModulesConfig := ModulesDto.ModulesConfig{
		Repositories.InitRepo(connection),
	}

	utilities := Dto.Utilities{
		Modules: Modules.Modules{
			UserModule: User.NewModules(ModulesConfig),
		},
	}
	newConfig.Routes = &Routes.Routes{
		Gin:        routesConfig,
		Controller: Controllers.InitControllerApi(utilities),
	}
	routes := newConfig.Routes.CollectRoutes()
	newConfig.HttpEngine.Run(routes)
}
