package main

import (
	"github.com/hadihammurabi/belajar-go-graphql/config"
	deliveryGraphql "github.com/hadihammurabi/belajar-go-graphql/internal/app/delivery/gql"

	_ "github.com/hadihammurabi/belajar-go-graphql/docs"

	"github.com/joho/godotenv"
)

// @title Belajar Go REST API
// @version 0.0.1
// @description Ini adalah projek untuk latihan REST API dengan Go
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	_ = godotenv.Load()

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ioc := NewIOC(conf)
	httpApp := deliveryGraphql.Init(ioc)

	forever := make(chan bool)
	go func() {
		httpApp.HTTP.Listen(conf.APP.Port)
	}()
	<-forever
}
