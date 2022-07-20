package vo

import (
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"net/url"
)

type OriginalUrl struct {
	value string
}

func NewOriginalUrl(value string) (originalUrlId OriginalUrl, err error) {
	originalUrlId = OriginalUrl{value: value}
	if err = originalUrlId.hasError(); err != nil {
		originalUrlId = OriginalUrl{}
	}
	return
}

func (u *OriginalUrl) hasError() error {
	if u.value == "" {
		return fmt.Errorf("%w", exception.ErrEmptyOriginalUrl)
	}
	if !IsUrl(u.value) {
		return fmt.Errorf("%w: %s", exception.ErrInvalidOriginalUrl, u.value)
	}
	return nil
}

func (u OriginalUrl) Value() string {
	return u.value
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
