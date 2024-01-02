package repository

import (
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/pkg/shared"
)

type TokenRepository interface {
	shared.CRUDRepository[*models.Token, models.Token]
}
