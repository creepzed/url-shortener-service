package getting

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllUrlShortenerQuery(t *testing.T) {
	t.Parallel()
	t.Run("given a valid request, the query is built", func(t *testing.T) {
		userId := randomvalues.RandomUserId()

		queryType := GetAllUrlShortenerQueryType

		qry := NewGetAllUrlShortenerQuery(userId)

		assert.Equal(t, qry.Type(), queryType)
		assert.Equal(t, qry.UserId(), userId)
	})
}
