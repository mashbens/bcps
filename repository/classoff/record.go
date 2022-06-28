package classoff

import (
	"github.com/mashbens/cps/business/classoff/entity"
)

type Offlineclass struct {
	ID          int    `gorm:"primary_key:auto_increment"`
	Classname   string `gorm:"type:varchar(100)" `
	Trainer     string `gorm:"type:varchar(100)" `
	Date        string `gorm:"type:varchar(100)" `
	Clock       string `gorm:"type:varchar(100)" `
	Description string `gorm:"type:varchar(100)" `
}

func (c *Offlineclass) toService() entity.Classoff {
	return entity.Classoff{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
	}
}

func fromService(c entity.Classoff) Offlineclass {
	return Offlineclass{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
	}
}

func toServiceList(data []Offlineclass) []entity.Classoff {
	a := []entity.Classoff{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
