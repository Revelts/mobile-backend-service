package Configs

import (
	"database/sql"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

type DbSqlConfigName string
type dbRedisConfigName string

const (
	// Database Connection Constant name
	databaseMain DbSqlConfigName   = "mainDB"
	redisMainDb  dbRedisConfigName = "mainDBRedis"
)

type DbConInterface interface {
	MySQLMainCon() *sql.DB
	setSqlConnection(conName DbSqlConfigName, con *sql.DB)
	RedisMainCon() *redis.Client
	setRedisConnection(conName dbRedisConfigName, con *redis.Client)
}

type DbCon struct {
	sql   map[DbSqlConfigName]*sql.DB
	redis map[dbRedisConfigName]*redis.Client
}

func (d DbCon) MySQLMainCon() *sql.DB {
	return d.sql[databaseMain]
}

func (d *DbCon) setSqlConnection(conName DbSqlConfigName, con *sql.DB) {
	if d.sql == nil {
		d.sql = make(map[DbSqlConfigName]*sql.DB)
	}
	d.sql[conName] = con
}

func (d *DbCon) setRedisConnection(conName dbRedisConfigName, con *redis.Client) {
	if d.redis == nil {
		d.redis = make(map[dbRedisConfigName]*redis.Client)
	}
	d.redis[conName] = con
}

func (d DbCon) RedisMainCon() *redis.Client {
	return d.redis[redisMainDb]
}
