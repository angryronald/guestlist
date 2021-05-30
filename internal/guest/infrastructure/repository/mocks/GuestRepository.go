// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	repository "github.com/angryronald/guestlist/internal/guest/infrastructure/repository"
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// GuestRepository is an autogenerated mock type for the GuestRepository type
type GuestRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, guest
func (_m *GuestRepository) Delete(ctx context.Context, guest *repository.Guest) error {
	ret := _m.Called(ctx, guest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *repository.Guest) error); ok {
		r0 = rf(ctx, guest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, isArrivedOnly
func (_m *GuestRepository) FindAll(ctx context.Context, isArrivedOnly bool) ([]*repository.Guest, error) {
	ret := _m.Called(ctx, isArrivedOnly)

	var r0 []*repository.Guest
	if rf, ok := ret.Get(0).(func(context.Context, bool) []*repository.Guest); ok {
		r0 = rf(ctx, isArrivedOnly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, isArrivedOnly)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, guestID
func (_m *GuestRepository) FindByID(ctx context.Context, guestID uuid.UUID) (*repository.Guest, error) {
	ret := _m.Called(ctx, guestID)

	var r0 *repository.Guest
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *repository.Guest); ok {
		r0 = rf(ctx, guestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, guestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: ctx, name
func (_m *GuestRepository) FindByName(ctx context.Context, name string) (*repository.Guest, error) {
	ret := _m.Called(ctx, name)

	var r0 *repository.Guest
	if rf, ok := ret.Get(0).(func(context.Context, string) *repository.Guest); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, guest
func (_m *GuestRepository) Insert(ctx context.Context, guest *repository.Guest) (*repository.Guest, error) {
	ret := _m.Called(ctx, guest)

	var r0 *repository.Guest
	if rf, ok := ret.Get(0).(func(context.Context, *repository.Guest) *repository.Guest); ok {
		r0 = rf(ctx, guest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *repository.Guest) error); ok {
		r1 = rf(ctx, guest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, guest
func (_m *GuestRepository) Update(ctx context.Context, guest *repository.Guest) (*repository.Guest, error) {
	ret := _m.Called(ctx, guest)

	var r0 *repository.Guest
	if rf, ok := ret.Get(0).(func(context.Context, *repository.Guest) *repository.Guest); ok {
		r0 = rf(ctx, guest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *repository.Guest) error); ok {
		r1 = rf(ctx, guest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
