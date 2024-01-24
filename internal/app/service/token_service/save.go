package tokenService

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/common/lib/app_errors"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
)

func (t *tokenService) Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*models.Token, error) {
	token, err := t.TokenRepository.GetById(ctx, data.Id)
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	if token == nil {
		token, err := t.TokenRepository.Create(ctx, &models.Token{Id: data.Id, Value: data.RefreshToken})
		if err != nil {
			return nil, appErrors.InternalServerError("")
		}
		return token, nil
	}

	token, err = t.TokenRepository.Update(ctx, data.Id, &models.Token{Id: data.Id, Value: data.RefreshToken})
	if err != nil {
		return nil, appErrors.InternalServerError("")
	}
	return token, nil
}
