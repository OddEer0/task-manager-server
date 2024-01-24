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

func (t *tokenRepository) DeleteByValue(ctx context.Context, value string) error {
	t.mock.Tokens = lo.Filter(t.mock.Tokens, func(item *models.Token, index int) bool {
		if item.Value == value {
			return false
		}
		return true
	})
	return nil
}

func (t *tokenRepository) Create(ctx context.Context, data *models.Token) (*models.Token, error) {
	token := models.Token{Id: data.Id, Value: data.Value}
	t.mock.Tokens = append(t.mock.Tokens, &token)
	return &token, nil
}

func (t *tokenRepository) GetById(ctx context.Context, id string) (*models.Token, error) {
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

func (t *tokenRepository) Update(ctx context.Context, id string, data *models.Token) (*models.Token, error) {
	for _, item := range t.mock.Tokens {
		if item.Id == id {
			item.Id = data.Id
			item.Value = data.Value
			return item, nil
		}
	}
	return nil, nil
}

func (t *tokenRepository) Delete(ctx context.Context, id string) error {
	t.mock.Tokens = lo.Filter(t.mock.Tokens, func(item *models.Token, index int) bool {
		if item.Id == id {
			return false
		}
		return true
	})
	return nil
}

func (t *tokenRepository) HasByValue(ctx context.Context, value string) (bool, error) {
	for _, token := range t.mock.Tokens {
		if token.Value == value {
			return true, nil
		}
	}
	return false, nil
}

func NewTokenRepository() repository.TokenRepository {
	return &tokenRepository{mock.NewMockToken()}
}
