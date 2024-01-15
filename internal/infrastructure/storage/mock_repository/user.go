package mock_repository

import (
	"context"
	"fmt"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	"github.com/OddEer0/task-manager-server/internal/presentation/mock"
	"github.com/samber/lo"
)

type userRepository struct {
	mock *mock.MockedUser
}

func (u *userRepository) Create(ctx context.Context, data *aggregate.UserAggregate) (*aggregate.UserAggregate, error) {
	newUser := models.User{
		Id:             data.User.Id,
		Nick:           data.User.Nick,
		Password:       data.User.Password,
		FirstName:      data.User.FirstName,
		LastName:       data.User.LastName,
		Email:          data.User.Email,
		ActivationLink: data.User.ActivationLink,
		Role:           data.User.Role,
	}
	u.mock.Users = append(u.mock.Users, &newUser)
	return &aggregate.UserAggregate{User: newUser}, nil
}

func (u *userRepository) GetById(ctx context.Context, id string) (*aggregate.UserAggregate, error) {
	var searched *models.User = nil
	for _, user := range u.mock.Users {
		if user.Id == id {
			searched = user
		}
	}
	if searched != nil {
		return &aggregate.UserAggregate{User: *searched}, nil
	} else {
		return nil, nil
	}
}

func (u *userRepository) Update(ctx context.Context, id string, data *aggregate.UserAggregate) (*aggregate.UserAggregate, error) {
	has := false
	lo.Map(u.mock.Users, func(user *models.User, i int) *models.User {
		if id == user.Id {
			has = true
			copyUser := data.User
			return &copyUser
		}
		return user
	})
	if has {
		return data, nil
	} else {
		return nil, fmt.Errorf("not found")
	}
}

func (u *userRepository) Delete(ctx context.Context, id string) error {
	has := false
	lo.Filter(u.mock.Users, func(user *models.User, index int) bool {
		if user.Id != id {
			return true
		}
		has = true
		return false
	})
	if has {
		return nil
	} else {
		return fmt.Errorf("not found")
	}
}

func (u *userRepository) HasUserByNick(ctx context.Context, nick string) (bool, error) {
	return lo.ContainsBy(u.mock.Users, func(item *models.User) bool {
		if item.Nick == nick {
			return true
		}
		return false
	}), nil
}

func (u *userRepository) HasUserByEmail(ctx context.Context, email string) (bool, error) {
	return lo.ContainsBy(u.mock.Users, func(item *models.User) bool {
		if item.Email.Value == email {
			return true
		}
		return false
	}), nil
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{mock.NewMockUser()}
}
