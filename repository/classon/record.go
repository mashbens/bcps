package classon

import (
	"github.com/mashbens/cps/business/classon/entity"
)

type ClassOnline struct {
	ID          int    `gorm:"primary_key:auto_increment" json:"-"`
	Classname   string `gorm:"type:varchar(100)" `
	Trainer     string `gorm:"type:varchar(100)" `
	Date        string `gorm:"type:varchar(100)" `
	Clock       string `gorm:"type:varchar(100)" `
	Description string `gorm:"type:varchar(100)" `
	AdminID     int
	Admin       entity.Admin `gorm:"foreignkey:AdminID" json:"-"`
}

func (c *ClassOnline) toService() entity.Classon {
	return entity.Classon{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
		AdminID:     c.AdminID,
		Admin:       c.Admin,
	}
}

func fromService(c entity.Classon) ClassOnline {
	return ClassOnline{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
		AdminID:     c.AdminID,
		Admin:       c.Admin,
	}
}

func toServiceList(data []ClassOnline) []entity.Classon {
	a := []entity.Classon{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
