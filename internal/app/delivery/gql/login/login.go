package login

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/entity"
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/service"
)

type LoginSchema struct {
	Service *service.Service
	Type    *graphql.Object
}

func New(service *service.Service) *LoginSchema {
	schema := graphql.NewObject(graphql.ObjectConfig{
		Name: "Login",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	return &LoginSchema{
		Service: service,
		Type:    schema,
	}
}

func (schema *LoginSchema) AsField() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewNonNull(schema.Type),
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "user's email",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "user's password",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			email, ok := p.Args["email"].(string)
			if !ok {
				return nil, errors.New("email required")
			}

			password, ok := p.Args["password"].(string)
			if !ok {
				return nil, errors.New("password required")
			}

			loginResult, err := schema.Service.Auth.Login(&entity.User{
				Email:    email,
				Password: password,
			})

			if err != nil {
				return nil, err
			}

			return map[string]string{
				"token": loginResult,
			}, nil
		},
	}
}
