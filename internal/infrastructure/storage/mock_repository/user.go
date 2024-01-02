package mock_repository

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
)

type userRepository struct {
	mock *mock.MockedUser
}

func (u userRepository) Create(ctx context.Context, data dto.CreateUserDto) (*aggregate.UserAggregate, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetById(ctx context.Context, id string) (*aggregate.UserAggregate, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update(ctx context.Context, data *aggregate.UserAggregate) (*aggregate.UserAggregate, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() repository.UserRepository {
	return userRepository{mock.NewMockUser()}
}
