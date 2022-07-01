package booking

import (
	"github.com/mashbens/cps/business/booking"
	"github.com/mashbens/cps/util"
)

func BookingRepoFactory(dbCon *util.DatabaseConnection) booking.BookingRepo {
	var bookingRepo booking.BookingRepo

	if dbCon.Driver == util.PostgreSQL {
		bookingRepo = NewBookingPostgresRepo(dbCon.PostgreSQL)
		dbCon.PostgreSQL.AutoMigrate(&Booking{})

	} else {
		panic("Database driver not supported")
	}

	return bookingRepo
}
