// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	mock "github.com/stretchr/testify/mock"
)

// mockRouteTablesClient is an autogenerated mock type for the routeTablesClient type
type mockRouteTablesClient struct {
	mock.Mock
}

// ListAll provides a mock function with given fields: options
func (_m *mockRouteTablesClient) ListAll(options *armnetwork.RouteTablesListAllOptions) routeTablesListAllPager {
	ret := _m.Called(options)

	var r0 routeTablesListAllPager
	if rf, ok := ret.Get(0).(func(*armnetwork.RouteTablesListAllOptions) routeTablesListAllPager); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(routeTablesListAllPager)
		}
	}

	return r0
}
