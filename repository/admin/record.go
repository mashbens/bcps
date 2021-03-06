package admin

import (
	"github.com/mashbens/cps/business/admin/entity"
)

type Admin struct {
	ID           int    `gorm:"primary_key:auto_increment" json:"-"`
	Name         string `gorm:"type:varchar(100)" json:"-"`
	Email        string `gorm:"type:varchar(100)" json:"-"`
	Phone        string `gorm:"type:varchar(100)" json:"-"`
	Password     string `gorm:"type:varchar(100)" json:"-"`
	SuperAdminID int
	SuperAdmin   entity.SuperAdmin `gorm:"foreignkey:SuperAdminID" json:"-"`
}

func (a *Admin) toService() entity.Admin {
	return entity.Admin{
		ID:           a.ID,
		Name:         a.Name,
		Email:        a.Email,
		Password:     a.Password,
		Phone:        a.Phone,
		SuperAdminID: a.SuperAdminID,
		SuperAdmin:   a.SuperAdmin,
	}
}

func fromService(a entity.Admin) Admin {
	return Admin{
		ID:           a.ID,
		Name:         a.Name,
		Email:        a.Email,
		Password:     a.Password,
		Phone:        a.Phone,
		SuperAdminID: a.SuperAdminID,
		SuperAdmin:   a.SuperAdmin,
	}
}
func toServiceList(data []Admin) []entity.Admin {
	a := []entity.Admin{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
