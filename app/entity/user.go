package entity

const UserTableName = "tbl_user"

type User struct {
	Id            int    `gorm:"column:id;type:INT(11);PRIMARY_KEY;AUTO_INCREMENT;NOT NULL;" json:"id"`
	FullName      string `gorm:"column:full_name;type:VARCHAR(255);NOT NULL;default:'';" json:"full_name"`
	Avatar        string `gorm:"column:avatar;type:VARCHAR(100);NOT NULL;default:'';" json:"avatar"`
	Email         string `gorm:"column:email;type:VARCHAR(20);NOT NULL;INDEX;UNIQUE;" json:"email"`
	Phone         string `gorm:"column:phone;type:VARCHAR(15);NOT NULL;INDEX;UNIQUE;" json:"phone"`
	Password      string `gorm:"column:password;type:VARCHAR(255);NOT NULL;default:'';" json:"password"`
	TimeLastLogin int    `gorm:"column:time_last_login;type:INT(11);NOT NULL;default:0;" json:"time_last_login"`
	TimeDuring    int    `gorm:"column:time_during;type:INT(11);NOT NULL;default:0;" json:"time_during"`
	Status        int    `gorm:"column:status;type:TINYINT(1);NOT NULL;INDEX;default:0;" json:"status"`
	CreatedAt     int    `gorm:"column:created_at;type:INT(11);NOT NULL;" json:"created_at"`
	UpdatedAt     int    `gorm:"column:updated_at;type:INT(11);NOT NULL;" json:"updated_at"`
}

func (User) TableName() string {
	return UserTableName
}
