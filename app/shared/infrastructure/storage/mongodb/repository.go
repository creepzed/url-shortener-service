package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type repositoryMongoDB struct {
	connection ConnectionMongo
	dbTimeout  time.Duration
}

func NewRepositoryMongoDB(connection ConnectionMongo, dbTimeout time.Duration) *repositoryMongoDB {
	return &repositoryMongoDB{
		connection: connection,
		dbTimeout:  dbTimeout,
	}
}

func (r *repositoryMongoDB) Create(ctx context.Context, aEntity interface{}) (err error) {
	ctxConnectionTimeout, connectionCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer connectionCancel()

	collection, err := r.connection.GetCollection(ctxConnectionTimeout)
	defer r.connection.Close(ctx)

	if err != nil {
		return err
	}

	ctxTimeout, InsertOneCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer InsertOneCancel()

	insert, err := collection.InsertOne(ctxTimeout, aEntity)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: %s", storage.ErrDuplicate, err.Error())
		}
		return err
	}

	if insert != nil {
		log.Debug("insert result: %s", insert.InsertedID)
	}

	return nil
}

func (r *repositoryMongoDB) FindById(ctx context.Context, filter map[string]interface{}) (interface{}, error) {
	ctxConnectionTimeout, connectionCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer connectionCancel()

	collection, err := r.connection.GetCollection(ctxConnectionTimeout)
	defer r.connection.Close(ctx)

	if err != nil {
		return nil, err
	}

	ctxTimeout, findByIdOneCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer findByIdOneCancel()

	result := collection.FindOne(ctxTimeout, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	doc := make(map[string]interface{}, 0)
	err = result.Decode(&doc)
	if err != nil {
		return nil, errors.New("error transforming data from mongo")
	}

	return doc, nil
}
