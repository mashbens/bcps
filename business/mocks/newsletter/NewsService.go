// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/mashbens/cps/business/newsletter/entity"
	mock "github.com/stretchr/testify/mock"
)

// NewsService is an autogenerated mock type for the NewsService type
type NewsService struct {
	mock.Mock
}

// DeleteNews provides a mock function with given fields: adminId, newsID
func (_m *NewsService) DeleteNews(adminId string, newsID string) error {
	ret := _m.Called(adminId, newsID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(adminId, newsID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllNews provides a mock function with given fields: search
func (_m *NewsService) FindAllNews(search string) []entity.News {
	ret := _m.Called(search)

	var r0 []entity.News
	if rf, ok := ret.Get(0).(func(string) []entity.News); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.News)
		}
	}

	return r0
}

// FindNewsByID provides a mock function with given fields: newsID
func (_m *NewsService) FindNewsByID(newsID string) (*entity.News, error) {
	ret := _m.Called(newsID)

	var r0 *entity.News
	if rf, ok := ret.Get(0).(func(string) *entity.News); ok {
		r0 = rf(newsID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(newsID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertNews provides a mock function with given fields: news
func (_m *NewsService) InsertNews(news entity.News) (*entity.News, error) {
	ret := _m.Called(news)

	var r0 *entity.News
	if rf, ok := ret.Get(0).(func(entity.News) *entity.News); ok {
		r0 = rf(news)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.News) error); ok {
		r1 = rf(news)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNews provides a mock function with given fields: news
func (_m *NewsService) UpdateNews(news entity.News) (*entity.News, error) {
	ret := _m.Called(news)

	var r0 *entity.News
	if rf, ok := ret.Get(0).(func(entity.News) *entity.News); ok {
		r0 = rf(news)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.News) error); ok {
		r1 = rf(news)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewNewsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewsService creates a new instance of NewsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewsService(t mockConstructorTestingTNewNewsService) *NewsService {
	mock := &NewsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
