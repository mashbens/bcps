// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/mashbens/cps/business/class/entity"
	mock "github.com/stretchr/testify/mock"
)

// ClassService is an autogenerated mock type for the ClassService type
type ClassService struct {
	mock.Mock
}

// DeleteClass provides a mock function with given fields: adminId, classID
func (_m *ClassService) DeleteClass(adminId string, classID string) error {
	ret := _m.Called(adminId, classID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(adminId, classID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllClass provides a mock function with given fields: search
func (_m *ClassService) FindAllClass(search string) []entity.Class {
	ret := _m.Called(search)

	var r0 []entity.Class
	if rf, ok := ret.Get(0).(func(string) []entity.Class); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Class)
		}
	}

	return r0
}

// FindAllClassBySerach provides a mock function with given fields: search
func (_m *ClassService) FindAllClassBySerach(search string) []entity.Class {
	ret := _m.Called(search)

	var r0 []entity.Class
	if rf, ok := ret.Get(0).(func(string) []entity.Class); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Class)
		}
	}

	return r0
}

// FindAllClassOff provides a mock function with given fields: search
func (_m *ClassService) FindAllClassOff(search string) []entity.Class {
	ret := _m.Called(search)

	var r0 []entity.Class
	if rf, ok := ret.Get(0).(func(string) []entity.Class); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Class)
		}
	}

	return r0
}

// FindAllClassOn provides a mock function with given fields: search
func (_m *ClassService) FindAllClassOn(search string) []entity.Class {
	ret := _m.Called(search)

	var r0 []entity.Class
	if rf, ok := ret.Get(0).(func(string) []entity.Class); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Class)
		}
	}

	return r0
}

// FindClassByID provides a mock function with given fields: classID
func (_m *ClassService) FindClassByID(classID string) (*entity.Class, error) {
	ret := _m.Called(classID)

	var r0 *entity.Class
	if rf, ok := ret.Get(0).(func(string) *entity.Class); ok {
		r0 = rf(classID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Class)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(classID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindClassOffByID provides a mock function with given fields: classID
func (_m *ClassService) FindClassOffByID(classID string) (*entity.Class, error) {
	ret := _m.Called(classID)

	var r0 *entity.Class
	if rf, ok := ret.Get(0).(func(string) *entity.Class); ok {
		r0 = rf(classID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Class)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(classID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindClassOnByID provides a mock function with given fields: classID
func (_m *ClassService) FindClassOnByID(classID string) (*entity.Class, error) {
	ret := _m.Called(classID)

	var r0 *entity.Class
	if rf, ok := ret.Get(0).(func(string) *entity.Class); ok {
		r0 = rf(classID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Class)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(classID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertClass provides a mock function with given fields: _a0
func (_m *ClassService) InsertClass(_a0 entity.Class) (*entity.Class, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Class
	if rf, ok := ret.Get(0).(func(entity.Class) *entity.Class); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Class)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Class) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClass provides a mock function with given fields: _a0
func (_m *ClassService) UpdateClass(_a0 entity.Class) (*entity.Class, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Class
	if rf, ok := ret.Get(0).(func(entity.Class) *entity.Class); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Class)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Class) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClassStatus provides a mock function with given fields: classID, status
func (_m *ClassService) UpdateClassStatus(classID string, status string) error {
	ret := _m.Called(classID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(classID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserBooked provides a mock function with given fields: classID
func (_m *ClassService) UpdateUserBooked(classID string) error {
	ret := _m.Called(classID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(classID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClassService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClassService creates a new instance of ClassService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClassService(t mockConstructorTestingTNewClassService) *ClassService {
	mock := &ClassService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
