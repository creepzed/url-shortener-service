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
	err := IsValidUrlId(u.value)
	if err != nil {
		return err
	}
	return nil
}

func IsValidUrlId(urlId string) error {
	pattern := "[0-9A-Za-z]$"
	if urlId == "" {
		return fmt.Errorf("%w", exception.ErrEmptyUrlId)
	}
	match, _ := regexp.MatchString(pattern, urlId)
	if !match {
		return fmt.Errorf("%w: %s", exception.ErrInvalidUrlId, urlId)
	}
	return nil
}

func (u UrlId) Value() string {
	return u.value
}
