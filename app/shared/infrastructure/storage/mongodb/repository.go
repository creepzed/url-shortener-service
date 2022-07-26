package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type repositoryMongoDB struct {
	connection ConnectionMongo
	dbTimeout  time.Duration
}

var (
	ErrTransform = errors.New("error transforming data from mongo")
)

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
		return r.errorHandler(err)
	}

	ctxTimeout, InsertOneCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer InsertOneCancel()

	insert, err := collection.InsertOne(ctxTimeout, aEntity)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return r.errorHandler(fmt.Errorf("%w: %s", storage.ErrDuplicate, err.Error()))
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
		return nil, r.errorHandler(err)
	}

	ctxTimeout, findByIdOneCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer findByIdOneCancel()

	result := collection.FindOne(ctxTimeout, filter)
	if result.Err() != nil {
		return nil, r.errorHandler(result.Err())
	}

	doc := make(map[string]interface{}, 0)
	err = result.Decode(&doc)
	if err != nil {
		return nil, r.errorHandler(err)
	}

	return doc, nil
}

func (r *repositoryMongoDB) Update(ctx context.Context, filter map[string]interface{}, anAggregate interface{}) (err error) {
	ctxConnectionTimeout, connectionCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer connectionCancel()

	collection, err := r.connection.GetCollection(ctxConnectionTimeout)
	defer r.connection.Close(ctx)

	if err != nil {
		return r.errorHandler(err)
	}
	doc := map[string]interface{}{"$set": anAggregate}
	opt := options.FindOneAndUpdate()

	cursor := collection.FindOneAndUpdate(ctx, filter, doc, opt)

	if cursor.Err() != nil {
		return r.errorHandler(err)
	}

	return nil
}

func (r *repositoryMongoDB) errorHandler(err error) error {
	var auxErr error
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		auxErr = fmt.Errorf("%w: %s", exception.ErrUrlNotFound, err.Error())
	case errors.Is(err, mongo.ErrInvalidIndexValue):
		auxErr = fmt.Errorf("%w: %s", storage.ErrUniqueValueIndex, err.Error())
	default:
		auxErr = fmt.Errorf("%w: %s", exception.ErrDataBase, err.Error())
	}
	return auxErr
}

func (r *repositoryMongoDB) Find(ctx context.Context, filter map[string]interface{}) (listAggregate []map[string]interface{}, err error) {
	ctxConnectionTimeout, connectionCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer connectionCancel()

	collection, err := r.connection.GetCollection(ctxConnectionTimeout)
	defer r.connection.Close(ctx)

	if err != nil {
		return []map[string]interface{}{}, r.errorHandler(err)
	}

	ctxTimeout, findByIdOneCancel := context.WithTimeout(ctx, r.dbTimeout)
	defer findByIdOneCancel()

	cursor, err := collection.Find(ctxTimeout, filter)
	if err != nil {
		return []map[string]interface{}{}, r.errorHandler(err)
	}

	var docs []map[string]interface{}
	if err = cursor.All(ctx, &docs); err != nil {
		return []map[string]interface{}{}, err
	}
	return docs, nil
}
