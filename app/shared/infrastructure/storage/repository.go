package storage

import (
	"context"
	"errors"
)

var (
	ErrDuplicate        = errors.New("the id is duplicated")
	ErrUniqueValueIndex = errors.New("invalid index value")
)

type Repository interface {
	Create(ctx context.Context, anAggregate interface{}) (err error)
	FindById(ctx context.Context, filter map[string]interface{}) (anAggregate interface{}, err error)
	Update(ctx context.Context, filter map[string]interface{}, anAggregate interface{}) (err error)
	Find(ctx context.Context, filter map[string]interface{}) (listAggregate []map[string]interface{}, err error)
}
