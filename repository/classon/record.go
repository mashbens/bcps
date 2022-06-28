package classon

import (
	"github.com/mashbens/cps/business/classon/entity"
)

type Onlineclass struct {
	ID          int    `gorm:"primary_key:auto_increment" json:"-"`
	Classname   string `gorm:"type:varchar(100)" `
	Trainer     string `gorm:"type:varchar(100)" `
	Date        string `gorm:"type:varchar(100)" `
	Clock       string `gorm:"type:varchar(100)" `
	Description string `gorm:"type:varchar(100)" `
	// AdminID     int
	// Admin       Admin `gorm:"foreignKey:Refer:Admin;joinForeignKey:AdminID"`
}

// type Admin struct {
// 	ID       int
// 	Name     string
// 	Password string
// 	Email    string
// 	Phone    string
// }

func (c *Onlineclass) toService() entity.Classon {
	return entity.Classon{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
		// AdminID:     c.AdminID,
		// Admin:       entity.Admin(c.Admin),
	}
}

func fromService(c entity.Classon) Onlineclass {
	return Onlineclass{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
		// AdminID:     c.AdminID,
		// Admin:       Admin(c.Admin),
	}
}

func toServiceList(data []Onlineclass) []entity.Classon {
	a := []entity.Classon{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
