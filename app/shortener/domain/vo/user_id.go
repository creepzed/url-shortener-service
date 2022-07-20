package vo

import (
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"net/mail"
)

type UserId struct {
	value string
}

func NewUserId(value string) (userId UserId, err error) {
	userId = UserId{value: value}
	if err = userId.hasError(); err != nil {
		userId = UserId{}
	}
	return
}

func (u *UserId) hasError() error {
	_, err := mail.ParseAddress(u.value)
	if err != nil {
		return fmt.Errorf("%w: %s", exception.ErrInvalidUserId, u.Value())
	}
	return nil
}

func (u UserId) Value() string {
	return u.value
}
