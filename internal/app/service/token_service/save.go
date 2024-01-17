package tokenService

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
)

func (t *tokenService) Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*models.Token, error) {
	return nil, nil
}
