package resp

import (
	booking "github.com/mashbens/cps/business/booking/entity"
)

type BookResp struct {
	ID    int   `json:"id"`
	User  User  `json:"user"`
	Class Class `json:"class"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type Class struct {
	ID          int    `json:"id"`
	Classname   string `json:"classname"`
	ClassType   string `json:"ClassType"`
	Trainer     string `json:"Trainer"`
	Date        string `json:"Date"`
	Clock       string `json:"Clock"`
	Description string `json:"Description"`
}

func FromService(c booking.Booking) BookResp {
	return BookResp{
		ID: int(c.ID),
		User: User{
			ID:    c.User.ID,
			Name:  c.User.Name,
			Email: c.User.Email,
			Phone: c.User.Phone,
		},
		Class: Class{
			ID:          c.Class.ID,
			Classname:   c.Class.Classname,
			ClassType:   c.Class.ClassType,
			Trainer:     c.Class.Trainer,
			Date:        c.Class.Date,
			Clock:       c.Class.Clock,
			Description: c.Class.Description,
		},
	}
}
