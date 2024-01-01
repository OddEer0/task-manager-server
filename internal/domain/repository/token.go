package repository

import (
	"task-manager-server/internal/domain/models"
	"task-manager-server/internal/pkg/shared"
)

type TokenRepository interface {
	shared.CRUDRepository[*models.Token, models.Token]
}
