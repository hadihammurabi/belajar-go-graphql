package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/delivery/gql/login"
)

func InitSchema(delivery *Delivery) *handler.Handler {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"login": login.New(delivery.Service).AsField(),
			},
		}),
	})
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return h
}
