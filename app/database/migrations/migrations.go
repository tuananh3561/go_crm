package migrations

import (
	"github.com/tuananh3561/go_crm/app/config"
	"github.com/tuananh3561/go_crm/app/entity"
	"log"
)

func Migrations(database config.Database) {
	err := database.MysqlAuth.Debug().AutoMigrate(
		&entity.User{},
		&entity.UserCountry{},
		&entity.UserRole{},
		&entity.UserPermission{},
		&entity.Role{},
		&entity.Permission{},
		&entity.PermissionRole{},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Migration successfully.")
}
