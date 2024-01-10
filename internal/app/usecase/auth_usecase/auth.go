package authUsecase

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/app/service/token_service"
	"github.com/OddEer0/task-manager-server/internal/app/service/user_service"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

type (
	AuthResult struct {
		User   *dto.ResponseUserDto
		Tokens tokenService.JwtTokens
	}

	AuthUseCase interface {
		Registration(ctx context.Context, data dto.RegistrationInputDto) (*AuthResult, error)
		Login(ctx context.Context, data dto.LoginInputDto) (*AuthResult, error)
		Logout(ctx context.Context, refreshToken string) error
		Refresh(ctx context.Context, refreshToken string) (*AuthResult, error)
	}

	authUseCase struct {
		repository.UserRepository
		UserService  userService.Service
		TokenService tokenService.Service
	}
)

func NewAuthUseCase(userService userService.Service, tokenService tokenService.Service) AuthUseCase {
	return &authUseCase{UserService: userService, TokenService: tokenService}
}
