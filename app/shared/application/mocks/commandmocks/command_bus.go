// Code generated by mockery v2.12.2. DO NOT EDIT.

package commandmocks

import (
	context "context"

	command "github.com/creepzed/url-shortener-service/app/shared/application/command"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// CommandBus is an autogenerated mock type for the CommandBus type
type CommandBus struct {
	mock.Mock
}

// Dispatch provides a mock function with given fields: _a0, _a1
func (_m *CommandBus) Dispatch(_a0 context.Context, _a1 command.Command) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.Command) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: _a0, _a1
func (_m *CommandBus) Register(_a0 command.Type, _a1 command.Handler) {
	_m.Called(_a0, _a1)
}

// NewCommandBus creates a new instance of CommandBus. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommandBus(t testing.TB) *CommandBus {
	mock := &CommandBus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
