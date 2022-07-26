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

		queryType := FindUrlShortenerQueryType

		metadata := Metadata{}
		qry := NewFindUrlShortenerQuery(urlId, metadata)

		assert.Equal(t, qry.Type(), queryType)
		assert.Equal(t, qry.UrlId(), urlId)
		assert.Equal(t, qry.Metadata(), metadata)
	})
}
