package resp

import (
	booking "github.com/mashbens/cps/business/booking/entity"
)

type BookRespSlice struct {
	UserID     int             `json:"user_id"`
	ClassSlice []booking.Class `json:"class"`
}

func FromServiceSlice(c booking.Booking) BookRespSlice {
	return BookRespSlice{
		UserID:     int(c.UserID),
		ClassSlice: c.ClassSlice,
	}
}

type BookResp struct {
	UserID int           `json:"user_id"`
	Class  booking.Class `json:"class"`
}

func FromService(c booking.Booking) BookResp {
	return BookResp{
		UserID: int(c.UserID),
		Class:  c.Class,
	}
}
