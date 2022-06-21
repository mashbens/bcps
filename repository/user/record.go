package user

import (
	"github.com/mashbens/cps/business/user/entity"
)

type User struct {
	ID             int    `gorm:"primary_key:auto_increment" json:"-"`
	Name           string `gorm:"type:varchar(100)" json:"-"`
	Email          string `gorm:"type:varchar(100);unique;" json:"-"`
	Phone          string `gorm:"type:varchar(100);unique;" json:"-"`
	Password       string `gorm:"type:varchar(100)" json:"-"`
	Member_expired string `gorm:"type:varchar(100)" json:"-"`
	Member_type    string `gorm:"type:varchar(100)" json:"-"`
}

func (u *User) toService() entity.User {
	return entity.User{
		ID:             int(u.ID),
		Name:           u.Name,
		Email:          u.Email,
		Phone:          u.Phone,
		Password:       u.Password,
		Member_expired: u.Member_expired,
		Member_type:    u.Member_type,
	}
}

func fromService(user entity.User) User {
	return User{
		ID:             int(user.ID),
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		Password:       user.Password,
		Member_expired: user.Member_expired,
		Member_type:    user.Member_type,
	}
}
