// Code generated by mockery v2.12.2. DO NOT EDIT.

package storagemocks

import (
	"github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// KeyGenerateService is an autogenerated mock type for the KeyGenerateService type
type KeyGenerateService struct {
	mock.Mock
}

// GetKey provides a mock function with given fields:
func (_m *KeyGenerateService) GetKey() (vo.UrlId, error) {
	ret := _m.Called()

	var r0 vo.UrlId
	if rf, ok := ret.Get(0).(func() vo.UrlId); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vo.UrlId)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKeyGenerateService creates a new instance of KeyGenerateService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewKeyGenerateService(t testing.TB) *KeyGenerateService {
	mock := &KeyGenerateService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
