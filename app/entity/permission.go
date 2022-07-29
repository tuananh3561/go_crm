package entity

const PermissionTableName = "tbl_permission"

type Permission struct {
	ID        string `gorm:"column:id;type:VARCHAR(20);PRIMARY_KEY;NOT NULL;" json:"id"`
	Name      string `gorm:"column:name;type:VARCHAR(255);NOT NULL;default:'';" json:"name"`
	Router    string `gorm:"column:router;type:VARCHAR(50);NOT NULL;INDEX;UNIQUE;default:'';" json:"router"`
	Note      string `gorm:"column:note;type:TEXT;" json:"note"`
	Type      int    `gorm:"column:type;type:TINYINT(1);NOT NULL;INDEX;default:0;" json:"type"`
	IsPublic  bool   `gorm:"column:is_public;type:BOOLEAN;NOT NULL;INDEX;default:false;" json:"is_public"`
	Status    int    `gorm:"column:status;type:TINYINT(1);NOT NULL;INDEX;default:0;" json:"status"`
	CreatedAt int    `gorm:"column:created_at;type:INT(11);NOT NULL;" json:"created_at"`
	UpdatedAt int    `gorm:"column:updated_at;type:INT(11);NOT NULL;" json:"updated_at"`
}

func (Permission) TableName() string {
	return PermissionTableName
}
