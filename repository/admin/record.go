package admin

import (
	"github.com/mashbens/cps/business/admin/entity"
)

type Admin struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"-"`
	Name     string `gorm:"type:varchar(100)" json:"-"`
	Email    string `gorm:"type:varchar(100);unique;" json:"-"`
	Phone    string `gorm:"type:varchar(100);unique;" json:"-"`
	Password string `gorm:"type:varchar(100)" json:"-"`
}

func (a *Admin) toService() entity.Admin {
	return entity.Admin{
		ID:       a.ID,
		Name:     a.Name,
		Email:    a.Email,
		Password: a.Password,
		Phone:    a.Phone,
	}
}

func fromService(a entity.Admin) Admin {
	return Admin{
		ID:       a.ID,
		Name:     a.Name,
		Email:    a.Email,
		Password: a.Password,
		Phone:    a.Phone,
	}
}
