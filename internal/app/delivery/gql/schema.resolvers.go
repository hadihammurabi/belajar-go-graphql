package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hadihammurabi/belajar-go-graphql/internal/app/delivery/gql/generated"
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/delivery/gql/model"
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/entity"
)

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.LoginResult, error) {
	loginResult, err := r.Delivery.Service.Auth.Login(&entity.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return &model.LoginResult{
		Token: loginResult,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
