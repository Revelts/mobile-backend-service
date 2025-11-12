package Configs

import (
	"bytes"
	"database/sql"
	"github.com/go-redis/redis"
	"log"
	"mobile-banking-service/Constants"
)

type connectionPool interface {
	MysqlConnectionPool(Connection interface{})
	RedisConnectionPool(Connection interface{})
}

func (env *database) MysqlConnectionPool(Connection interface{}) {
	var buffer bytes.Buffer
	Con := Connection.(*sql.DB)
	buffer.WriteString(env.Username + ":" + env.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(env.Host + ":" + env.Port + ")/")
	buffer.WriteString(env.Name)
	connection_string := buffer.String()
	Connection, err := sql.Open(Constants.MYSQL, connection_string)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	Connection.(*sql.DB).SetMaxOpenConns(env.Maximum_connection)
	*Con = *Connection.(*sql.DB)
	err = Con.Ping()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
		return
	}
	return
}

func (env *database) RedisConnectionPool(Connection interface{}) {
	var buffer bytes.Buffer
	Con := Connection.(*redis.Client)
	buffer.WriteString(env.Host + ":" + env.Port)
	connectionString := buffer.String()

	Connection = redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: env.Password,
		DB:       0,
	})

	*Con = *Connection.(*redis.Client)

	_, err := Con.Ping().Result()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
		return
	}
}
