package migrations

import (
	"github.com/tuananh3561/go_crm/app/entity"
	"gorm.io/gorm"
	"log"
)

func Migrations(db *gorm.DB) {
	err := db.Debug().AutoMigrate(
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
