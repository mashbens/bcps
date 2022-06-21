package migrate

import "gorm.io/gorm"

type User struct {
	ID             int    `gorm:"primary_key:auto_increment" json:"-"`
	Name           string `gorm:"type:varchar(100)" json:"-"`
	Email          string `gorm:"type:varchar(100);unique;" json:"-"`
	Phone          string `gorm:"type:varchar(100);unique;" json:"-"`
	Password       string `gorm:"type:varchar(100)" json:"-"`
	Member_expired string `gorm:"type:varchar(100)" json:"-"`
	Member_type    string `gorm:"type:varchar(100)" json:"-"`
	*gorm.Model
}
