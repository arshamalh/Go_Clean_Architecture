// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "coryptex.com/bot/vip-signal/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// AdminRepository is an autogenerated mock type for the AdminRepository type
type AdminRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *AdminRepository) Add(_a0 context.Context, _a1 entities.Admin) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, entities.Admin) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entities.Admin) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTID provides a mock function with given fields: _a0, _a1
func (_m *AdminRepository) GetByTID(_a0 context.Context, _a1 string) (*entities.Admin, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entities.Admin
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.Admin); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Admin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
