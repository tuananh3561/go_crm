package entity

const UserPermissionTableName = "tbl_user_permission"

type UserPermission struct {
	UserId       int    `gorm:"column:user_id;type:INT(11);PRIMARY_KEY;NOT NULL;INDEX;" json:"user_id"`
	PermissionId string `gorm:"column:permission_id;type:VARCHAR(20);PRIMARY_KEY;NOT NULL;INDEX;" json:"permission_id"`
	Type         int    `gorm:"column:type;type:TINYINT(1);NOT NULL;INDEX;default:0;" json:"type"`
	CreatedAt    int    `gorm:"column:created_at;type:INT(11);NOT NULL;" json:"created_at"`
	UpdatedAt    int    `gorm:"column:updated_at;type:INT(11);NOT NULL;" json:"updated_at"`
}

func (UserPermission) TableName() string {
	return UserPermissionTableName
}
