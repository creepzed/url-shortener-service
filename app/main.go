package main

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/bus/inmemory"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/rest"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage/mongodb"
	"github.com/creepzed/url-shortener-service/app/shortener/application/creating"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/mongo"
	"os"
	"time"
)

const (
	host      = "localhost"
	port      = 8080
	dbTimeOut = 5 * time.Second
)

var (
	mongoDBURI        = os.Getenv("MONGODB_URI")
	mongoDBName       = os.Getenv("MONGODB_DATABASE")
	mongoDBCollection = os.Getenv("MONGODB_COLLECTION")
)

func main() {
	server := rest.New()

	mongodbConn := mongodb.NewMongoDBConnection(mongoDBURI, mongoDBName, mongoDBCollection)
	dbUrlShortener := mongo.NewUrlShortenerRepositoryMongo(mongodbConn, dbTimeOut)

	commandBusInMemory := inmemory.NewCommandBusMemory()
	queryBusInMemory := inmemory.NewQueryBusMemory()
	eventBusInMemory := inmemory.NewEventBusInMemory()

	createService := creating.NewCreateApplicationService(dbUrlShortener, eventBusInMemory)

	createCommandHandler := creating.NewCreateUrlShortenerCommandHandler(createService)

	commandBusInMemory.Register(creating.CreateUrlShortenerCommandType, createCommandHandler)

	controllers.NewUrlShortenerController(server, commandBusInMemory, queryBusInMemory)

	server.Logger.Fatal(server.StartServer(rest.Setup(host, port)))
}
