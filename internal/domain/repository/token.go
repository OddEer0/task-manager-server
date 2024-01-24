package repository

import (
	"context"

	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/pkg/shared"
)

type TokenRepository interface {
	shared.CRUDRepository[*models.Token, *models.Token]
	DeleteByValue(ctx context.Context, value string) error
	HasByValue(ctx context.Context, value string) (bool, error)
}
