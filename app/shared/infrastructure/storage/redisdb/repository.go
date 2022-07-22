package redisdb

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"time"
)

type RepositoryRedis struct {
	connection ConnectionRedis
	dbTimeout  time.Duration
}

func NewRepositoryRedis(connection ConnectionRedis, dbTimeout time.Duration) *RepositoryRedis {
	return &RepositoryRedis{
		connection: connection,
		dbTimeout:  dbTimeout,
	}
}

func (r RepositoryRedis) Set(ctx context.Context, key string, value interface{}) error {
	ctxConnectionTimeout, connectionCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer connectionCancel()

	client, err := r.connection.GetConnection(ctxConnectionTimeout)
	defer r.connection.Close()
	if err != nil {
		return err
	}

	ctxTimeout, Cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer Cancel()

	str := utils.EntityToJson(value)
	err = client.WithContext(ctxTimeout).Set(key, str, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RepositoryRedis) Get(ctx context.Context, key string) (interface{}, error) {
	ctxConnectionTimeout, connectionCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer connectionCancel()

	client, err := r.connection.GetConnection(ctxConnectionTimeout)
	defer r.connection.Close()
	if err != nil {
		return nil, err
	}
	ctxTimeout, Cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer Cancel()

	result := client.WithContext(ctxTimeout).Get(key)
	if result.Err() != nil {
		return nil, result.Err()
	}

	str, err := result.Result()
	if err != nil {
		return nil, err
	}

	value := new(interface{})

	utils.JsonToEntity(str, &value)

	return value, err
}

func (r RepositoryRedis) Remove(ctx context.Context, key string) error {
	ctxConnectionTimeout, connectionCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer connectionCancel()

	client, err := r.connection.GetConnection(ctxConnectionTimeout)
	defer r.connection.Close()
	if err != nil {
		return err
	}
	ctxTimeout, Cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer Cancel()

	result := client.WithContext(ctxTimeout).Del(key)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
