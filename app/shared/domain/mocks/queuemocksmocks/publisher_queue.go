// Code generated by mockery v2.12.2. DO NOT EDIT.

package queuemocks

import (
	context "context"

	queue "github.com/creepzed/url-shortener-service/app/shared/domain/queue"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// PublisherQueue is an autogenerated mock type for the PublisherQueue type
type PublisherQueue struct {
	mock.Mock
}

// Close provides a mock function with given fields: topic
func (_m *PublisherQueue) Close(topic string) error {
	ret := _m.Called(topic)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(topic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: ctx, topic, messageData
func (_m *PublisherQueue) Publish(ctx context.Context, topic string, messageData queue.MessageData) error {
	ret := _m.Called(ctx, topic, messageData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, queue.MessageData) error); ok {
		r0 = rf(ctx, topic, messageData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPublisherQueue creates a new instance of PublisherQueue. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPublisherQueue(t testing.TB) *PublisherQueue {
	mock := &PublisherQueue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
