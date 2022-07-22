package redisdb

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/go-redis/redis"
	"strconv"
)

type ConnectionRedis interface {
	GetConnection(ctx context.Context) (*redis.Client, error)
	Close()
}

type DbConnection struct {
	addr     *string
	password *string
	db       *int
	client   *redis.Client
}

func NewRedisDBConnection(addr string, password string, db string) *DbConnection {
	if addr == "" {
		log.Fatal("error address is not valid")
	}
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		log.Fatal("error db is not valid")
	}

	return &DbConnection{
		addr:     &addr,
		password: &password,
		db:       &dbInt,
	}
}

func (d *DbConnection) GetConnection(ctx context.Context) (*redis.Client, error) {
	d.client = redis.NewClient(&redis.Options{
		Addr:     *d.addr,
		Password: *d.password,
		DB:       *d.db,
	})
	pong, err := d.client.Ping().Result()
	if err != nil {
		return nil, err
	}
	log.Info("Redis connection: %s", pong)
	return d.client, nil
}

func (d DbConnection) Close() {
	d.client.Close()
	d.client = nil
}
