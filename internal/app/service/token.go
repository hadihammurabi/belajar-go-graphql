package service

import (
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/entity"
	"github.com/hadihammurabi/belajar-go-graphql/internal/app/repository"
	"github.com/sarulabs/di"

	uuid "github.com/satori/go.uuid"
)

// TokenService struct
type TokenService struct {
	repo *repository.Repository
}

// NewTokenService func
func NewTokenService(ioc di.Container) entity.TokenService {
	repo := getRepository(ioc)
	return &TokenService{
		repo,
	}
}

// Create func
func (u TokenService) Create(token *entity.Token) (*entity.Token, error) {
	return u.repo.Token.Create(token)
}

// FindByUserID func
func (u TokenService) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	return u.repo.Token.FindByUserID(id)
}

// FindByToken func
func (u TokenService) FindByToken(token string) (*entity.Token, error) {
	return u.repo.Token.FindByToken(token)
}
