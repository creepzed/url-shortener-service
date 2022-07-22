package finding

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindUrlShortenerQuery(t *testing.T) {
	t.Parallel()
	t.Run("given a valid request, the query is built", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()

		commandType := FindUrlShortenerQueryType

		qry := NewFindUrlShortenerQuery(urlId)

		assert.Equal(t, qry.Type(), commandType)
		assert.Equal(t, qry.UrlId(), urlId)
	})
}