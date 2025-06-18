package main

import (
	catdb "CatTracker/internal/cat/db"
	cathandler "CatTracker/internal/cat/handler"
	catservice "CatTracker/internal/cat/service"
	eventdb "CatTracker/internal/event/db"
	eventhandler "CatTracker/internal/event/handler"
	eventservice "CatTracker/internal/event/service"
	"CatTracker/internal/general/config"
	"CatTracker/internal/general/db"
	"CatTracker/internal/general/kafka"
	"context"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.AppConfig
	kafkaClient, err := kafka.NewKafkaClient(&conf.Kafka)
	if err != nil {
		panic(err)
	}

	pgClient, err := db.NewPostgresClient(context.Background(), &conf.Db)
	if err != nil {
		panic(err)
	}

	catRepo := catdb.NewCatRepository(pgClient)
	eventRepo := eventdb.NewEventRepository(pgClient)

	catService := catservice.NewCatService(kafkaClient, catRepo)
	catHandler := cathandler.NewCatHandler(catService)

	eventService := eventservice.NewEventService(kafkaClient, eventRepo)
	eventHandler := eventhandler.NewEventHandler(eventService)

	ginEngine := gin.Default()

	catHandler.RegisterRoutes(ginEngine)
	eventHandler.RegisterRoutes(ginEngine)

	err = ginEngine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
