package reporting

import (
	"context"
	queuemocks "github.com/creepzed/url-shortener-service/app/shared/domain/mocks/queuemocksmocks"
	"github.com/creepzed/url-shortener-service/app/shortener/application/finding"
	"github.com/creepzed/url-shortener-service/app/shortener/application/mocks/servicemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestReportWrapApplicationService(t *testing.T) {
	t.Parallel()

	t.Run("given an error in find service the error is stored", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		query := finding.NewFindUrlShortenerQuery(urlId, nil)

		serviceFindMock := servicemocks.NewFindApplicationService(t)
		serviceFindMock.
			On("Do", context.Background(), mock.AnythingOfType("finding.FindUrlShortenerQuery")).
			Return(nil, exception.ErrDataBase)

		producerMock := queuemocks.NewPublisherQueue(t)

		service := NewReportApplicationService(serviceFindMock, producerMock, "topic")
		result, err := service.Do(context.Background(), query)

		assert.Error(t, err)
		assert.Equal(t, nil, result)

	})

	//t.Run("given a correct answer is recorded", func(t *testing.T) {
	//	urlId := randomvalues.RandomUrlId()
	//	query := finding.NewFindUrlShortenerQuery(urlId, nil)
	//
	//	serviceFindMock := servicemocks.NewFindApplicationService(t)
	//	serviceFindMock.
	//		On("Do", context.Background(), mock.AnythingOfType("finding.FindUrlShortenerQuery")).
	//		Return(nil, exception.ErrDataBase)
	//
	//	producerMock := queuemocks.NewPublisherQueue(t)
	//	producerMock.
	//		On("Publish", context.Background(), mock.AnythingOfType("string"), mock.AnythingOfType("queue.MessageData")).
	//		Return(nil)
	//
	//	service := NewReportApplicationService(serviceFindMock, producerMock, "topic")
	//
	//	result, err := service.Do(context.Background(), query)
	//
	//	assert.Error(t, err)
	//	assert.Equal(t, nil, result)
	//
	//})
}
