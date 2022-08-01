package mysqldb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigMysql struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection(config ConfigMysql) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database mysqldb")
		return nil
	}
	fmt.Println("Connected to Mysql!")
	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection mysqldb")
	}
	err = dbSQL.Close()
	if err != nil {
		panic("Failed to close database database mysqldb")
	}
}
