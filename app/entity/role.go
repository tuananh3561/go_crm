package entity

const RoleTableName = "tbl_role"

type Role struct {
	Id        string `gorm:"column:id;type:VARCHAR(20);PRIMARY_KEY;NOT NULL;" json:"id"`
	Name      string `gorm:"column:name;type:VARCHAR(255);NOT NULL;default:'';" json:"name"`
	ParentId  string `gorm:"column:parent_id;type:VARCHAR(20);INDEX;" json:"parent_id"`
	Status    int    `gorm:"column:status;type:TINYINT(1);NOT NULL;INDEX;default:0;" json:"status"`
	CreatedAt int    `gorm:"column:created_at;type:INT(11);NOT NULL;" json:"created_at"`
	UpdatedAt int    `gorm:"column:updated_at;type:INT(11);NOT NULL;" json:"updated_at"`
}

func (Role) TableName() string {
	return RoleTableName
}

const RoleStatusActive = 1
const RoleStatusInactive = 2
