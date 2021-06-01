package gql

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hadihammurabi/belajar-go-graphql/config"
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/service"
	"github.com/sarulabs/di"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Delivery struct
type Delivery struct {
	HTTP        *fiber.App
	Middlewares func(int) fiber.Handler
	Service     *service.Service
	Validator   *config.Validator
	Config      *config.Config
}

// Init func
func Init(ioc di.Container) *Delivery {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[\"${time}\", \"${method}\", \"${path}\", \"${status}\", \"${ip}\", \"${latency}\"]\n",
	}))
	app.Use(recover.New())
	app.Use(cors.New())

	service := ioc.Get("service").(*service.Service)
	conf := ioc.Get("config").(*config.Config)

	delivery := &Delivery{
		HTTP:      app,
		Service:   service,
		Config:    conf,
		Validator: config.NewValidator(),
	}

	h := InitSchema(delivery)
	delivery.HTTP.Use("/gql", adaptor.HTTPHandler(h))
	return delivery
}
