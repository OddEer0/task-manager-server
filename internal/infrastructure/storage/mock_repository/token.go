package mock_repository

import (
	"context"

	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
	"github.com/samber/lo"
)

type tokenRepository struct {
	mock *mock.MockedToken
}

func (t tokenRepository) Create(ctx context.Context, data *models.Token) (*models.Token, error) {
	t.mock.Tokens = append(t.mock.Tokens, data)
	return data, nil
}

func (t tokenRepository) GetById(ctx context.Context, id string) (*models.Token, error) {
	token, ok := lo.Find(t.mock.Tokens, func(item *models.Token) bool {
		if item.Id == id {
			return true
		}
		return false
	})
	if !ok {
		return nil, nil
	}
	return token, nil
}

func (t tokenRepository) Update(ctx context.Context, id string, data *models.Token) (*models.Token, error) {
	for _, item := range t.mock.Tokens {
		if item.Id == id {
			item.Id = data.Id
			item.Value = data.Value
			return item, nil
		}
	}
	return nil, nil
}

func (t tokenRepository) Delete(ctx context.Context, id string) error {
	t.mock.Tokens = lo.Filter(t.mock.Tokens, func(item *models.Token, index int) bool {
		if item.Id == id {
			return false
		}
		return true
	})
	return nil
}

func NewTokenRepository() repository.TokenRepository {
	return &tokenRepository{mock.NewMockToken()}
}
