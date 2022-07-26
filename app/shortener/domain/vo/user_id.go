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
	err := IsValidUserId(u.value)
	if err != nil {
		return err
	}
	return nil
}

func (u UserId) Value() string {
	return u.value
}

func IsValidUserId(userId string) error {
	_, err := mail.ParseAddress(userId)
	if err != nil {
		return fmt.Errorf("%w: %s", exception.ErrInvalidUserId, userId)
	}
	return nil
}
