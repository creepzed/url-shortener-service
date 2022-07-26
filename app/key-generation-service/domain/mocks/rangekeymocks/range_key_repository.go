// Code generated by mockery v2.12.2. DO NOT EDIT.

package rangekeymocks

import (
	domain "github.com/creepzed/url-shortener-service/app/key-generation-service/domain"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RangeKeyRepository is an autogenerated mock type for the RangeKeyRepository type
type RangeKeyRepository struct {
	mock.Mock
}

// GetRange provides a mock function with given fields:
func (_m *RangeKeyRepository) GetRange() (*domain.RangeKey, error) {
	ret := _m.Called()

	var r0 *domain.RangeKey
	if rf, ok := ret.Get(0).(func() *domain.RangeKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.RangeKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRangeKeyRepository creates a new instance of RangeKeyRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRangeKeyRepository(t testing.TB) *RangeKeyRepository {
	mock := &RangeKeyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
