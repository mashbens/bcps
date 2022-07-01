package request

import "github.com/mashbens/cps/business/booking/entity"

type CreateBookingRequest struct {
	UserID  int `json:"user_id"`
	ClassID int `json:"class_id"`
}

func NewCreateBookingReq(req CreateBookingRequest) entity.Booking {
	return entity.Booking{
		UserID:  req.UserID,
		ClassID: req.ClassID,
	}
}
