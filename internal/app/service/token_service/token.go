package tokenService

import (
	"context"

	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/dgrijalva/jwt-go"
)

type (
	JwtTokens struct {
		AccessToken  string
		RefreshToken string
	}

	CustomClaims struct {
		JwtUserData dto.GenerateTokenDto
		jwt.StandardClaims
	}

	Service interface {
		Generate(data dto.GenerateTokenDto) (*JwtTokens, error)
		Save(ctx context.Context, data dto.SaveTokenDto) (*models.Token, error)
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
