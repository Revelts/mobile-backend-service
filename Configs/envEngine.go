package Configs

import (
	"database/sql"
	"io/ioutil"
	"log"
	"mobile-banking-service/Constants"
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v3"
)

func GetEnvironment(env string) Config {
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), Constants.ENVIRONMENT_PATH+env+".yml")
	_, err := os.Stat(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
		return nil
	}
	content, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	var config envFile = content
	return config
}

func (e envFile) LoadConfig() *ConfigSetting {

	var config Environment

	err := yaml.Unmarshal([]byte(string(e)), &config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	if config.App.Debug == false {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("Environment Configs load successfully!")
	return &ConfigSetting{&config, nil, &config.App, &config}
}

func (e *Environment) BuildConnection() DbConInterface {
	var connectionPool connectionPool = &database{}
	var dbCon DbCon
	for i := 0; i < len(e.Databases); i++ {
		connectionPool = &e.Databases[i]
		switch e.Databases[i].Engine {
		case Constants.MYSQL:
			con := sql.DB{}
			log.Println("ENGINE " + Constants.MYSQL + " start....")
			connectionPool.MysqlConnectionPool(&con)
			dbCon.setSqlConnection(DbSqlConfigName(e.Databases[i].Connection), &con)
		case Constants.REDIS:
			con := redis.Client{}
			log.Println("ENGINE " + Constants.REDIS + " start....")
			connectionPool.RedisConnectionPool(&con)
			dbCon.setRedisConnection(dbRedisConfigName(e.Databases[i].Connection), &con)
		}

	}
	return &dbCon
}

func (app *app) Run(route *gin.Engine) {
	//run with https or http
	if app.Service == "https" {
		app.runWithHttps(route)
	}
	app.runWithHttp(route)
}
func (app *app) runWithHttp(route *gin.Engine) {

	log.Println("Http Service running ....")
	address := app.Host + ":" + app.Port
	err := route.Run(address)
	if err != nil {
		//panic error
		panic(err)
	}
}

func (app *app) runWithHttps(route *gin.Engine) {

	log.Println("Https Service running ....")
	address := app.Host + ":" + app.Port
	// Setting Directory Path Key and Pem File
	_, filename, _, _ := runtime.Caller(1)
	filepathKey := path.Join(path.Dir(filename), "../Infrastructures/certificate/"+app.Pem_key)
	filepathCert := path.Join(path.Dir(filename), "../Infrastructures/certificate/"+app.Certificate)
	// Setting Listen TLS
	err := route.RunTLS(address, filepathCert, filepathKey)
	if err != nil {
		log.Println(err.Error())
		//panic error
		panic(err)
		return
	}
}
