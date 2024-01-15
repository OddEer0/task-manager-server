package mock_repository

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
)

type tokenRepository struct {
	mock *mock.MockedToken
}

func (t tokenRepository) Create(ctx context.Context, data *models.Token) (*models.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) GetById(ctx context.Context, id string) (*models.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) Update(ctx context.Context, id string, data *models.Token) (*models.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTokenRepository() repository.TokenRepository {
	return &tokenRepository{}
}
