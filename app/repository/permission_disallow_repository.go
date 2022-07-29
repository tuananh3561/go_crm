package repository

import "gorm.io/gorm"

type PermissionDisallowRepository interface {
}

type permissionDisallowConnection struct {
	connection *gorm.DB
}

func NewPermissionDisallowRepository(db *gorm.DB) PermissionDisallowRepository {
	return &permissionDisallowConnection{
		connection: db,
	}
}
