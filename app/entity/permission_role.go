package entity

const PermissionRoleTableName = "tbl_permission_role"

type PermissionRole struct {
	RoleId       string `gorm:"column:role_id;type:VARCHAR(20);PRIMARY_KEY;NOT NULL;INDEX;" json:"role_id"`
	PermissionId string `gorm:"column:permission_id;type:VARCHAR(20);PRIMARY_KEY;NOT NULL;INDEX;" json:"permission_id"`
	Rule         string `gorm:"column:rule;type:TEXT;" json:"rule"`
	CreatedAt    int    `gorm:"column:created_at;type:INT(11);NOT NULL;" json:"created_at"`
	UpdatedAt    int    `gorm:"column:updated_at;type:INT(11);NOT NULL;" json:"updated_at"`
}

func (PermissionRole) TableName() string {
	return PermissionRoleTableName
}
