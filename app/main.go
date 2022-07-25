package main

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/bus/inmemory"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/rest"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage/mongodb"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage/redisdb"
	"github.com/creepzed/url-shortener-service/app/shortener/application/creating"
	"github.com/creepzed/url-shortener-service/app/shortener/application/finding"
	"github.com/creepzed/url-shortener-service/app/shortener/application/reporting"
	"github.com/creepzed/url-shortener-service/app/shortener/application/updating"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/transformer"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/queue/eventbus"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/queue/kafka/common"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/queue/kafka/producer"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/cache"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/mongo"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/redis"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

const (
	host      = ""
	port      = 8080
	dbTimeOut = 5 * time.Second
)

var (
	mongoDBURI        = os.Getenv("MONGODB_URI")
	mongoDBName       = os.Getenv("MONGODB_DATABASE")
	mongoDBCollection = os.Getenv("MONGODB_COLLECTION")

	redisAddr     = os.Getenv("REDIS_ADDR")
	redisPassword = os.Getenv("REDIS_PASSWORD")
	redisDB       = os.Getenv("REDIS_DB")

	kafkaUsername = os.Getenv("KAFKA_USERNAME")
	kafkaPassword = os.Getenv("KAFKA_PASSWORD")
	kafkaBrokers  = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

	statisticsTopic = os.Getenv("KAFKA_STATISTICS_TOPIC")
	shortenerEvent  = os.Getenv("KAFKA_SHORTENER_EVENT_TOPIC")
)

func main() {
	server := rest.New()

	mongodbConn := mongodb.NewMongoDBConnection(mongoDBURI, mongoDBName, mongoDBCollection)
	repositoryMongo := mongo.NewUrlShortenerRepositoryMongo(mongodbConn, dbTimeOut)

	redisConn := redisdb.NewRedisDBConnection(redisAddr, redisPassword, redisDB)
	repositoryRedis := redis.NewUrlShortenerRepositoryRedis(redisConn, dbTimeOut)

	repositoryCache := cache.NewCache(repositoryMongo, repositoryRedis)

	producerQueue := producer.NewKafkaPublisher(common.GetDialer(kafkaUsername, kafkaPassword), kafkaBrokers...)

	commandBusInMemory := inmemory.NewCommandBusMemory()
	queryBusInMemory := inmemory.NewQueryBusMemory()
	eventBusInMemory := eventbus.NewEventBusKafka(producerQueue, shortenerEvent)

	createService := creating.NewCreateApplicationService(repositoryMongo, eventBusInMemory)

	createCommandHandler := creating.NewCreateUrlShortenerCommandHandler(createService)

	commandBusInMemory.Register(creating.CreateUrlShortenerCommandType, createCommandHandler)

	transform := transformer.NewTransformer()

	findService := finding.NewFindApplicationService(repositoryCache, transform)

	reportWrapService := reporting.NewReportApplicationService(findService, producerQueue, statisticsTopic)

	ctxObserver, cancelObserver := context.WithCancel(context.Background())
	defer cancelObserver()
	go reportWrapService.Observer(ctxObserver)

	findQueryHandler := finding.NewFindUrlShortenerQueryHandler(reportWrapService)

	queryBusInMemory.Register(finding.FindUrlShortenerQueryType, findQueryHandler)

	updateService := updating.NewUpdateApplicationService(repositoryMongo, eventBusInMemory)
	updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
	commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

	controllers.NewUrlShortenerController(server, commandBusInMemory, queryBusInMemory)

	go func() {
		if err := server.StartServer(rest.Setup(host, port)); err != http.ErrServerClosed {
			server.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}

}
