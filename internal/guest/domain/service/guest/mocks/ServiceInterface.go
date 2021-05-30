// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	public "github.com/angryronald/guestlist/internal/guest/public"

	uuid "github.com/google/uuid"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// CreateGuest provides a mock function with given fields: ctx, _a1
func (_m *ServiceInterface) CreateGuest(ctx context.Context, _a1 *public.Guest) (*public.Guest, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *public.Guest
	if rf, ok := ret.Get(0).(func(context.Context, *public.Guest) *public.Guest); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*public.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *public.Guest) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGuest provides a mock function with given fields: ctx, guestID
func (_m *ServiceInterface) DeleteGuest(ctx context.Context, guestID uuid.UUID) error {
	ret := _m.Called(ctx, guestID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, guestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAvailableSpace provides a mock function with given fields: ctx
func (_m *ServiceInterface) GetAvailableSpace(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGuest provides a mock function with given fields: ctx, guestID
func (_m *ServiceInterface) GetGuest(ctx context.Context, guestID uuid.UUID) (*public.Guest, error) {
	ret := _m.Called(ctx, guestID)

	var r0 *public.Guest
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *public.Guest); ok {
		r0 = rf(ctx, guestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*public.Guest)
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

// GetGuestByName provides a mock function with given fields: ctx, name
func (_m *ServiceInterface) GetGuestByName(ctx context.Context, name string) (*public.Guest, error) {
	ret := _m.Called(ctx, name)

	var r0 *public.Guest
	if rf, ok := ret.Get(0).(func(context.Context, string) *public.Guest); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*public.Guest)
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

// ListGuests provides a mock function with given fields: ctx, isArrivedOnly
func (_m *ServiceInterface) ListGuests(ctx context.Context, isArrivedOnly bool) ([]*public.Guest, error) {
	ret := _m.Called(ctx, isArrivedOnly)

	var r0 []*public.Guest
	if rf, ok := ret.Get(0).(func(context.Context, bool) []*public.Guest); ok {
		r0 = rf(ctx, isArrivedOnly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*public.Guest)
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

// UpdateGuest provides a mock function with given fields: ctx, _a1
func (_m *ServiceInterface) UpdateGuest(ctx context.Context, _a1 *public.Guest) (*public.Guest, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *public.Guest
	if rf, ok := ret.Get(0).(func(context.Context, *public.Guest) *public.Guest); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*public.Guest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *public.Guest) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
