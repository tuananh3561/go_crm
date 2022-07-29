package repository

import (
	"github.com/tuananh3561/go_crm/app/dto"
	"github.com/tuananh3561/go_crm/app/entity"
	"github.com/tuananh3561/go_crm/app/helper"
	"gorm.io/gorm"
)

type RoleRepository interface {
	RolesByParams(params dto.RoleSearchDTO) ([]entity.Role, error)
	CountRoleByParams(params dto.RoleSearchDTO) (int64, error)
	FindRoleById(id string) (*entity.Role, error)
	CreateRole(role entity.Role) (*entity.Role, error)
	UpdateRole(role entity.Role) error
}

type roleConnection struct {
	connection *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleConnection{
		connection: db,
	}
}

func (db roleConnection) RolesByParams(params dto.RoleSearchDTO) ([]entity.Role, error) {
	var roles []entity.Role

	result := db.connection
	result = queryRolesByParams(result, params)
	if params.Limit != 0 {
		offset := (params.Page - 1) * params.Limit
		result = result.Limit(params.Limit).Offset(offset)
	}
	err := result.Distinct().Order("created_at DESC").Find(&roles).Error
	return roles, err
}

func (db roleConnection) CountRoleByParams(params dto.RoleSearchDTO) (int64, error) {
	var total int64
	result := db.connection.Model(entity.Role{})
	result = queryRolesByParams(result, params)
	err := result.Distinct().Count(&total).Error
	return total, err
}

func queryRolesByParams(result *gorm.DB, params dto.RoleSearchDTO) *gorm.DB {
	if params.RoleId != "" {
		result = result.Where("id = ?", params.RoleId)
	}
	if len(params.RoleIds) > 0 {
		result = result.Where("id IN ?", params.RoleIds)
	}
	if params.Name != "" {
		result = getQueryLike(result, "name", params.Name)
	}
	if params.ParentId != "" {
		result = result.Where("parent_id = ?", params.ParentId)
	}
	if params.Status != 0 {
		result = result.Where("status = ?", params.Status)
	}
	return result
}

func (db roleConnection) FindRoleById(id string) (*entity.Role, error) {
	var role entity.Role
	err := db.connection.
		Where("id = ?", id).
		Find(&role).Error
	return &role, err
}

func (db roleConnection) CreateRole(role entity.Role) (*entity.Role, error) {
	err := CreateRole(db.connection, &role)
	return &role, err
}

func (db roleConnection) UpdateRole(role entity.Role) error {
	err := UpdateRole(db.connection, &role)
	return err
}

func CreateRole(tx *gorm.DB, role *entity.Role) error {
	if role.UpdatedAt == 0 {
		role.UpdatedAt = helper.TimeNow()
	}
	if role.UpdatedAt == 0 {
		role.CreatedAt = helper.TimeNow()
	}
	err := tx.Create(&role).Error
	return err
}

func UpdateRole(tx *gorm.DB, role *entity.Role) error {
	if role.UpdatedAt == 0 {
		role.UpdatedAt = helper.TimeNow()
	}
	err := tx.Where("id  = ?", role.Id).
		Updates(role).Error
	return err
}
