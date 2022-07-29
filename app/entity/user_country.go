package entity

const UserCountryTableName = "tbl_user_country"

type UserCountry struct {
	UserId      int    `gorm:"column:user_id;type:INT(11);PRIMARY_KEY;NOT NULL;INDEX;" json:"user_id"`
	CountryCode string `gorm:"column:country_code;type:VARCHAR(5);PRIMARY_KEY;NOT NULL;INDEX;" json:"country_code"`
	CreatedAt   int    `gorm:"column:created_at;type:INT(11);NOT NULL;" json:"created_at"`
	UpdatedAt   int    `gorm:"column:updated_at;type:INT(11);NOT NULL;" json:"updated_at"`
}

func (UserCountry) TableName() string {
	return UserCountryTableName
}
