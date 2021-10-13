// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	context "context"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/network/armnetwork"

	mock "github.com/stretchr/testify/mock"
)

// mockPublicIPAddressesListAllPager is an autogenerated mock type for the publicIPAddressesListAllPager type
type mockPublicIPAddressesListAllPager struct {
	mock.Mock
}

// Err provides a mock function with given fields:
func (_m *mockPublicIPAddressesListAllPager) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NextPage provides a mock function with given fields: ctx
func (_m *mockPublicIPAddressesListAllPager) NextPage(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PageResponse provides a mock function with given fields:
func (_m *mockPublicIPAddressesListAllPager) PageResponse() armnetwork.PublicIPAddressesListAllResponse {
	ret := _m.Called()

	var r0 armnetwork.PublicIPAddressesListAllResponse
	if rf, ok := ret.Get(0).(func() armnetwork.PublicIPAddressesListAllResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(armnetwork.PublicIPAddressesListAllResponse)
	}

	return r0
}
