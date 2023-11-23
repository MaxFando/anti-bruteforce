// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	network "github.com/MaxFando/anti-bruteforce/internal/domain/network"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, prefix, mask
func (_m *Store) Add(ctx context.Context, prefix string, mask string) error {
	ret := _m.Called(ctx, prefix, mask)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, prefix, mask)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx
func (_m *Store) Get(ctx context.Context) ([]network.IpNetwork, error) {
	ret := _m.Called(ctx)

	var r0 []network.IpNetwork
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]network.IpNetwork, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []network.IpNetwork); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]network.IpNetwork)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, prefix, mask
func (_m *Store) Remove(ctx context.Context, prefix string, mask string) error {
	ret := _m.Called(ctx, prefix, mask)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, prefix, mask)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}