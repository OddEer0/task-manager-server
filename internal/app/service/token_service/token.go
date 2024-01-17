package tokenService

import (
	"context"

	appDto "github.com/OddEer0/task-manager-server/internal/app/app_dto"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/dgrijalva/jwt-go"
)

type (
	JwtTokens struct {
		AccessToken  string
		RefreshToken string
	}

	CustomClaims struct {
		JwtUserData appDto.GenerateTokenServiceDto
		jwt.StandardClaims
	}

	Service interface {
		Generate(data appDto.GenerateTokenServiceDto) (*JwtTokens, error)
		Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*models.Token, error)
		Delete(ctx context.Context, id string) error
	}

	tokenService struct {
		repository.TokenRepository
	}
)

func New(tokenRepo repository.TokenRepository) Service {
	return &tokenService{
		TokenRepository: tokenRepo,
	}
}
