// Code generated by mockery v2.12.2. DO NOT EDIT.

package storagemocks

import (
	context "context"

	domain "github.com/creepzed/url-shortener-service/app/shortener/domain"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	vo "github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

// GetAllByUserIdRepository is an autogenerated mock type for the GetAllByUserIdRepository type
type GetAllByUserIdRepository struct {
	mock.Mock
}

// GetAllByUserId provides a mock function with given fields: ctx, userId
func (_m *GetAllByUserIdRepository) GetAllByUserId(ctx context.Context, userId vo.UserId) ([]domain.UrlShortener, error) {
	ret := _m.Called(ctx, userId)

	var r0 []domain.UrlShortener
	if rf, ok := ret.Get(0).(func(context.Context, vo.UserId) []domain.UrlShortener); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UrlShortener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, vo.UserId) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGetAllByUserIdRepository creates a new instance of GetAllByUserIdRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetAllByUserIdRepository(t testing.TB) *GetAllByUserIdRepository {
	mock := &GetAllByUserIdRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}