// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	cart "ecommerce/features/cart"

	mock "github.com/stretchr/testify/mock"
)

// CartData is an autogenerated mock type for the CartData type
type CartData struct {
	mock.Mock
}

// AddToCart provides a mock function with given fields: userID, productID, add
func (_m *CartData) AddToCart(userID uint, productID uint, add cart.Core) (cart.Core, error) {
	ret := _m.Called(userID, productID, add)

	var r0 cart.Core
	if rf, ok := ret.Get(0).(func(uint, uint, cart.Core) cart.Core); ok {
		r0 = rf(userID, productID, add)
	} else {
		r0 = ret.Get(0).(cart.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, cart.Core) error); ok {
		r1 = rf(userID, productID, add)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CartList provides a mock function with given fields: userID
func (_m *CartData) CartList(userID uint) ([]cart.Core, error) {
	ret := _m.Called(userID)

	var r0 []cart.Core
	if rf, ok := ret.Get(0).(func(uint) []cart.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID, cartID
func (_m *CartData) Delete(userID uint, cartID uint) error {
	ret := _m.Called(userID, cartID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateQty provides a mock function with given fields: userID, cartID, quantity
func (_m *CartData) UpdateQty(userID uint, cartID uint, quantity int) (cart.Core, error) {
	ret := _m.Called(userID, cartID, quantity)

	var r0 cart.Core
	if rf, ok := ret.Get(0).(func(uint, uint, int) cart.Core); ok {
		r0 = rf(userID, cartID, quantity)
	} else {
		r0 = ret.Get(0).(cart.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, int) error); ok {
		r1 = rf(userID, cartID, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCartData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartData creates a new instance of CartData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartData(t mockConstructorTestingTNewCartData) *CartData {
	mock := &CartData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
