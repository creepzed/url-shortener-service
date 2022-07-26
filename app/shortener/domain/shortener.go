package domain

import (
	eventBase "github.com/creepzed/url-shortener-service/app/shared/application/event"
	vo2 "github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/event"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type UrlShortener struct {
	urlId       vo2.UrlId
	urlEnabled  vo.UrlEnabled
	originalUrl vo.OriginalUrl
	userId      vo.UserId

	events []eventBase.Event
}

func NewUrlShortener(urlId vo2.UrlId, originalUrl vo.OriginalUrl, userId vo.UserId) (urlShortener UrlShortener) {

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

func LoadUrlShortener(urlId vo2.UrlId, urlEnabled vo.UrlEnabled, originalUrl vo.OriginalUrl, userId vo.UserId) UrlShortener {
	urlShortener := UrlShortener{
		urlId:       urlId,
		urlEnabled:  urlEnabled,
		originalUrl: originalUrl,
		userId:      userId,
	}
	return urlShortener
}

func (s UrlShortener) UrlId() vo2.UrlId {
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

func (s *UrlShortener) PullEvents() []eventBase.Event {
	eventAux := s.events
	//s.events = []event.Event{}
	return eventAux
}

func (s *UrlShortener) IsChanged() bool {
	return len(s.events) > 0
}

func (s *UrlShortener) Update(isEnabled vo.UrlEnabled, originalUrl vo.OriginalUrl, userId vo.UserId) {
	isChanged := false
	if isEnabled.Value() != s.urlEnabled.Value() {
		s.urlEnabled = isEnabled
		isChanged = true
	}

	if originalUrl.Value() != s.originalUrl.Value() {
		s.originalUrl = originalUrl
		isChanged = true
	}

	if userId.Value() != s.userId.Value() {
		s.userId = userId
		isChanged = true
	}

	if isChanged {
		s.Record(event.NewShortenerUpdatedEvent(
			s.UrlId().Value(),
			s.IsEnabled().Value(),
			s.OriginalUrl().Value(),
			s.UserId().Value(),
		))
	}
}
