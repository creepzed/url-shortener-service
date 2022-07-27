package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	dbTimeout = 30000 * time.Millisecond
)

var (
	ErrConnection = errors.New("error connecting to mongodb")
	ErrPing       = errors.New("error ping to mongodb")
)

type ConnectionMongo interface {
	GetCollection(ctx context.Context) (*mongo.Collection, error)
	Close(ctx context.Context)
}

type DbConnection struct {
	uri            string
	dataName       string
	collectionName string
	client         *mongo.Client
}

func NewMongoDBConnection(uri, dataName, collectionName string) *DbConnection {
	if dataName == "" {
		log.Fatal("error databaseName is not valid")
	}

	if collectionName == "" {
		log.Fatal("error collectionNames is not valid")
	}

	if uri == "" {
		log.Fatal("error uri connection is not valid")
	}

	return &DbConnection{
		uri:            uri,
		dataName:       dataName,
		collectionName: collectionName,
		client:         nil,
	}
}

func (d *DbConnection) GetCollection(ctx context.Context) (*mongo.Collection, error) {
	var err error
	if d.client == nil {
		d.client, err = getClient(ctx, d.uri)
		if err != nil {
			return nil, err
		}
	}

	if !d.isAvailable(ctx, d.client, dbTimeout) {
		return nil, ErrPing
	}

	return d.client.Database(d.dataName).Collection(d.collectionName), nil
}

func getClient(ctx context.Context, uri string) (*mongo.Client, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions).
		SetServerSelectionTimeout(dbTimeout)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.WithError(err).Fatal(ErrConnection.Error())
		return nil, fmt.Errorf("%w: %s", ErrConnection, err.Error())
	}
	return client, nil
}

func (d *DbConnection) isAvailable(ctx context.Context, client *mongo.Client, timeout time.Duration) bool {

	err := client.Ping(ctx, nil)
	if err != nil {
		log.WithError(err).Warn(ErrPing.Error())
		return false
	}
	return true
}

func (d *DbConnection) Close(ctx context.Context) {
	d.client.Disconnect(ctx)
	d.client = nil
}
