package booking

import (
	"github.com/mashbens/cps/business/booking/entity"
)

type Booking struct {
	ID     int
	UserID int
	// User    entity.User `gorm:"foreignkey:UserID" json:"-"`
	ClassID int
	// Class   entity.Class `gorm:"foreignkey:UserID" json:"-"`
}

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
