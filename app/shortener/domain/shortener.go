package domain

import (
	eventBase "github.com/creepzed/url-shortener-service/app/shared/application/event"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/event"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type UrlShortener struct {
	urlId       vo.UrlId
	urlEnabled  vo.UrlEnabled
	originalUrl vo.OriginalUrl
	userId      vo.UserId

	events []eventBase.Event
}

func NewUrlShortener(urlId vo.UrlId, originalUrl vo.OriginalUrl, userId vo.UserId) (urlShortener UrlShortener) {

	urlShortener = UrlShortener{
		urlId:       urlId,
		urlEnabled:  vo.NewUrlEnabled(vo.Enabled),
		originalUrl: originalUrl,
		userId:      userId,
	}

	urlShortener.Record(event.NewShortenerCreatedEvent(
		urlShortener.UrlId().Value(),
		urlShortener.IsEnabled().Value(),
		urlShortener.OriginalUrl().Value(),
		urlShortener.UserId().Value(),
	))

	return
}

func LoadUrlShortener(urlId vo.UrlId, urlEnabled vo.UrlEnabled, originalUrl vo.OriginalUrl, userId vo.UserId) UrlShortener {
	urlShortener := UrlShortener{
		urlId:       urlId,
		urlEnabled:  urlEnabled,
		originalUrl: originalUrl,
		userId:      userId,
	}
	return urlShortener
}

func (s UrlShortener) UrlId() vo.UrlId {
	return s.urlId
}

func (s UrlShortener) IsEnabled() vo.UrlEnabled {
	return s.urlEnabled
}

func (s UrlShortener) OriginalUrl() vo.OriginalUrl {
	return s.originalUrl
}

func (s UrlShortener) UserId() vo.UserId {
	return s.userId
}

func (s *UrlShortener) Record(event eventBase.Event) {
	s.events = append(s.events, event)
}

func (s UrlShortener) PullEvents() []eventBase.Event {
	eventAux := s.events
	//s.events = []event.Event{}
	return eventAux
}
