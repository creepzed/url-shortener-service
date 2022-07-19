package storage

import (
	"context"
	"errors"
)

var (
	ErrDuplicate = errors.New("the id is duplicated")
)

type Repository interface {
	Create(ctx context.Context, aEntity interface{}) (err error)
	FindById(ctx context.Context, filter map[string]interface{}) (aEntity interface{}, err error)
}
