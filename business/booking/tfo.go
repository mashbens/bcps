package booking

import (
	booking "github.com/mashbens/cps/business/booking/entity"
	_class "github.com/mashbens/cps/business/class/entity"
	_user "github.com/mashbens/cps/business/user/entity"
)

func UserToBoo(data _user.User) booking.User {
	return booking.User{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Phone: data.Phone,
	}
}

func ClassToBoo(c _class.Class) booking.Class {
	return booking.Class{
		ID:          c.ID,
		Classname:   c.Classname,
		Trainer:     c.Trainer,
		Date:        c.Date,
		Clock:       c.Clock,
		Description: c.Description,
		ClassType:   c.ClassType,
	}
}
