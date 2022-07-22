package exception

import "errors"

var (
	ErrUrlIdDuplicate = errors.New("error UrlId duplicate")
	ErrEmptyUrlId     = errors.New("the field UrlId can not be empty")
	ErrInvalidUrlId   = errors.New("the field UrlId is invalid")

	ErrInvalidUserId = errors.New("the field UserId is invalid")

	ErrEmptyOriginalUrl   = errors.New("the field OriginalUrl can not be empty")
	ErrInvalidOriginalUrl = errors.New("the field OriginalUrl is invalid")

	ErrDataBase = errors.New("error saving to database")
	ErrEventBus = errors.New("error saving to eventbus")

	ErrUrlNotFound = errors.New("url not found")

	ErrTransforming = errors.New("error transforming data")
)
