package class

import (
	"github.com/mashbens/cps/business/class/entity"
)

type Class struct {
	ID          int    `gorm:"primary_key:auto_increment" json:"-"`
	Classname   string `gorm:"type:varchar(100)" `
	Trainer     string `gorm:"type:varchar(100)" `
	Date        string `gorm:"type:varchar(100)" `
	Clock       string `gorm:"type:varchar(100)" `
	Description string `gorm:"type:varchar(100)" `
	ClassType   string `gorm:"type:varchar(100)" `
	Capacity    int
	UserBooked  int
	Duration    int
	Img         string
	Status      string `gorm:"type:varchar(100)" `
	AdminID     int
	Admin       Admin `gorm:"ForeignKey:AdminID"`
}

type Admin struct {
	ID       int
	Name     string
	Password string
	Email    string
	Phone    string
}

func (c *Class) toService() entity.Class {
	return entity.Class{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
		ClassType:   c.ClassType,
		Status:      c.Status,
		Capacity:    c.Capacity,
		UserBooked:  c.UserBooked,
		Duration:    c.Duration,
		Img:         c.Img,
		AdminID:     c.AdminID,
		Admin:       entity.Admin(c.Admin),
	}
}

func fromService(c entity.Class) Class {
	return Class{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
		ClassType:   c.ClassType,
		Status:      c.Status,
		Capacity:    c.Capacity,
		UserBooked:  c.UserBooked,
		Duration:    c.Duration,
		Img:         c.Img,
		AdminID:     c.AdminID,
		Admin:       Admin(c.Admin),
	}
}

func toServiceList(data []Class) []entity.Class {
	a := []entity.Class{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
