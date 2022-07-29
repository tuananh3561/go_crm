package repository

import (
	"gorm.io/gorm"
)

type UserRoleRepository interface {
}

type userRoleConnection struct {
	connection *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleConnection{
		connection: db,
	}
}
