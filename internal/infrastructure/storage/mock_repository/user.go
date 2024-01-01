package mock_repository

import (
	"context"
	"task-manager-server/internal/domain/aggregate"
	"task-manager-server/internal/domain/repository"
	"task-manager-server/internal/presentation/dto"
	"task-manager-server/internal/presentation/mock"
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
