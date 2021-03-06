// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/mashbens/cps/business/booking/entity"
	mock "github.com/stretchr/testify/mock"
)

// BookingService is an autogenerated mock type for the BookingService type
type BookingService struct {
	mock.Mock
}

// GetSchedule provides a mock function with given fields: userID
func (_m *BookingService) GetSchedule(userID string) (*entity.Booking, error) {
	ret := _m.Called(userID)

	var r0 *entity.Booking
	if rf, ok := ret.Get(0).(func(string) *entity.Booking); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBooking provides a mock function with given fields: _a0
func (_m *BookingService) InsertBooking(_a0 entity.Booking) (*entity.Booking, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Booking
	if rf, ok := ret.Get(0).(func(entity.Booking) *entity.Booking); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Booking) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookingService interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookingService creates a new instance of BookingService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookingService(t mockConstructorTestingTNewBookingService) *BookingService {
	mock := &BookingService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
