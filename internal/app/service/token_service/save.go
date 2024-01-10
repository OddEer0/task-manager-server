package tokenService

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

func (t *tokenService) Save(ctx context.Context, data dto.SaveTokenDto) (*models.Token, error) {
	return nil, nil
}
