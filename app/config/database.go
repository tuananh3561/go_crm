package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/tuananh3561/go_crm/app/database/db/mongodb"
	"github.com/tuananh3561/go_crm/app/database/db/mysqldb"
	"github.com/tuananh3561/go_crm/app/database/db/redisdb"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"strconv"
)

type Database struct {
	MysqlAuth *gorm.DB
	Redis     *redis.Client
	Mongo     *mongo.Client
}

func getConfigMysqlAuth() mysqldb.ConfigMysql {
	config := mysqldb.ConfigMysql{}

	config.Host = GetEnv("DB_HOST", "127.0.0.1")
	config.Port = GetEnv("DB_PORT", "3306")
	config.Database = GetEnv("DB_NAME", "edu_auth")
	config.User = GetEnv("DB_USER", "")
	config.Password = GetEnv("DB_PASS", "")

	return config
}

func getConfigRedis() redisdb.ConfigRedis {
	config := redisdb.ConfigRedis{}

	config.Host = GetEnv("REDIS_HOST", "127.0.0.1")
	config.Port = GetEnv("REDIS_PORT", "6379")
	config.Database = 0
	config.Password = GetEnv("REDIS_PASSWORD", "")

	databaseRedis, err := strconv.Atoi(GetEnv("REDIS_DB", "0"))
	if err == nil {
		config.Database = databaseRedis
	}

	return config
}

func getConfigMongodb() string {
	mongoDb := GetEnv("MONGODB_URL", "mongodb://early:abcd1234@34.87.186.176:27017/edu_backend")

	return mongoDb
}

func SetupDatabaseConnection() Database {
	configMysqlAuth := getConfigMysqlAuth()
	mysqlAuth := mysqldb.SetupDatabaseConnection(configMysqlAuth)

	configRedis := getConfigRedis()
	redis := redisdb.ConnectionClient(configRedis)

	configMongo := getConfigMongodb()
	mongoDb := mongodb.DBInstance(configMongo)

	db := Database{
		MysqlAuth: mysqlAuth,
		Redis:     redis,
		Mongo:     mongoDb,
	}
	return db
}

func CloseDatabaseConnection(db Database) {
	if db.MysqlAuth != nil {
		mysqldb.CloseDatabaseConnection(db.MysqlAuth)
	}
}
