package booking

import (
	"github.com/mashbens/cps/business/booking/entity"
)

type Booking struct {
	ID     int
	UserID int
	// User    entity.User `gorm:"foreignkey:UserID" json:"-"`
	ClassID int
	// 	Class   Class
}

// type Class struct {
// 	ID        int
// 	Classname string
// }

func (b *Booking) toService() entity.Booking {
	return entity.Booking{
		ID:     b.ID,
		UserID: b.UserID,
		// User:    b.User,
		ClassID: b.ClassID,
		// Class:   b.Class,
	}
}

func fromService(b entity.Booking) Booking {
	return Booking{
		ID:     b.ID,
		UserID: b.UserID,
		// User:    b.User,
		ClassID: b.ClassID,
		// Class:   b.Class,
	}
}

func toServiceList(data []Booking) []entity.Booking {
	a := []entity.Booking{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
