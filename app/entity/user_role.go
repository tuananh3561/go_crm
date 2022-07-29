package entity

const UserRoleTableName = "tbl_user_role"

type UserRole struct {
	UserId    int    `gorm:"column:user_id;type:INT(11);PRIMARY_KEY;NOT NULL;INDEX;" json:"user_id"`
	RoleId    string `gorm:"column:role_id;type:VARCHAR(20);PRIMARY_KEY;NOT NULL;INDEX;" json:"role_id"`
	CreatedAt int    `gorm:"column:created_at;type:INT(11);NOT NULL;" json:"created_at"`
	UpdatedAt int    `gorm:"column:updated_at;type:INT(11);NOT NULL;" json:"updated_at"`
}

func (UserRole) TableName() string {
	return UserRoleTableName
}
