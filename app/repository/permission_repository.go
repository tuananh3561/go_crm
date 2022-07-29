package repository

import "gorm.io/gorm"

type PermissionRepository interface {
}

type permissionConnection struct {
	connection *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionConnection{
		connection: db,
	}
}
