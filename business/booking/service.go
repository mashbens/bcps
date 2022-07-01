package booking

import (
	"errors"
	"strconv"

	"github.com/mashbens/cps/business/booking/entity"
	classService "github.com/mashbens/cps/business/class"
	userService "github.com/mashbens/cps/business/user"
)

type BookingRepo interface {
	InsertBooking(booking entity.Booking) (entity.Booking, error)
	// GetSchedule(userID string) (entity.Booking, error)
	// FindBookingByID(bookingID int) (entity.Booking, error)
}

type BookingService interface {
	InsertBooking(booking entity.Booking) (*entity.Booking, error)
	// GetSchedule(userID string) (*entity.Booking, error)
	// FindBookingByID(bookingID int) (*entity.Booking, error)
}

type bookingService struct {
	bookingRepo  BookingRepo
	userService  userService.UserService
	classService classService.ClassService
}

func NewBookingService(
	bookingRepo BookingRepo,
	userService userService.UserService,
	classService classService.ClassService,
) BookingService {
	return &bookingService{
		bookingRepo:  bookingRepo,
		userService:  userService,
		classService: classService,
	}
}

func (c *bookingService) InsertBooking(booking entity.Booking) (*entity.Booking, error) {
	// find user
	user, err := c.userService.FindUserByID(strconv.Itoa(booking.UserID))
	if err != nil {
		return nil, errors.New("User not found")
	}
	_ = user

	// find class

	// class, err := c.classService.FindClassByID(strconv.Itoa(booking.ClassID))
	// if err != nil {
	// 	return nil, errors.New("Class not found")
	// }
	// _ = class

	// ++ user booked
	// upt := c.classService.UpdateUserBooked(strconv.Itoa(booking.ClassID))
	// _ = upt
	// log.Println(upt)

	// book, err := c.bookingRepo.InsertBooking(booking)
	// if err != nil {
	// 	return nil, err
	// }

	// book.User = UserToBoo(*user)
	// book.Class = ClassToBoo(*class)

	// var tempBooking entity.Booking

	class, err := c.classService.FindClassByID(strconv.Itoa(booking.ClassID))
	if err != nil {
		return nil, errors.New("Class not found")
	}
	_ = class

	var tempStatus entity.Booking
	if class.UserBooked == class.Capacity-1 {
		status := "Full Booked"
		tempStatus.Class.Status = status
	} else {
		status := "Available"
		tempStatus.Class.Status = status
	}

	booking.Class.Status = tempStatus.Class.Status

	// update class status
	err = c.classService.UpdateClassStatus(strconv.Itoa(class.ID), booking.Class.Status)
	// ++ user booked
	upt := c.classService.UpdateUserBooked(strconv.Itoa(booking.ClassID))
	_ = upt

	book, err := c.bookingRepo.InsertBooking(booking)
	if err != nil {
		return nil, err
	}

	book.Class.UserBooked++
	book.Class = ClassToBoo(*class)
	book.User = UserToBoo(*user)

	return &book, nil

}

// func (c *bookingService) FindBookingByID(bookingID int) (*entity.Booking, error) {
// 	booking, err := c.bookingRepo.FindBookingByID(bookingID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &booking, nil
// }
