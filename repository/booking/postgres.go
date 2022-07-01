package booking

import (
	"github.com/mashbens/cps/business/booking"
	"github.com/mashbens/cps/business/booking/entity"
	"gorm.io/gorm"
)

type BookingPostgresRepo struct {
	db *gorm.DB
}

func NewBookingPostgresRepo(db *gorm.DB) booking.BookingRepo {
	return &BookingPostgresRepo{
		db: db,
	}
}

func (c *BookingPostgresRepo) InsertBooking(booking entity.Booking) (entity.Booking, error) {
	record := fromService(booking)
	res := c.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
