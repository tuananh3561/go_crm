package repository

import "gorm.io/gorm"

type PermissionRoleRepository interface {
}

type permissionRoleConnection struct {
	connection *gorm.DB
}

func NewPermissionRoleRepository(db *gorm.DB) PermissionRoleRepository {
	return &permissionRoleConnection{
		connection: db,
	}
}
