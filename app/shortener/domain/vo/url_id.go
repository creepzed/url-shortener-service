package vo

import (
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"regexp"
)

type UrlId struct {
	value string
}

func NewUrlId(value string) (urlId UrlId, err error) {
	urlId = UrlId{value: value}
	if err = urlId.hasError(); err != nil {
		urlId = UrlId{}
	}
	return
}

func (u *UrlId) hasError() error {
	pattern := "[0-9A-Za-z]$"
	if u.value == "" {
		return fmt.Errorf("%w", exception.ErrEmptyUrlId)
	}
	match, _ := regexp.MatchString(pattern, u.value)
	if !match {
		return fmt.Errorf("%w: %s", exception.ErrInvalidUrlId, u.value)
	}
	return nil
}

func (u UrlId) Value() string {
	return u.value
}
