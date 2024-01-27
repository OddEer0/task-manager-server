package mock_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	inMemDb "github.com/OddEer0/task-manager-server/internal/infrastructure/storage/in_mem_db"
	"github.com/samber/lo"
)

type userRepository struct {
	db *inMemDb.InMemDb
}

func (u *userRepository) Create(ctx context.Context, data *aggregate.UserAggregate) (*aggregate.UserAggregate, error) {
	has := lo.ContainsBy(u.db.Users, func(item *models.User) bool {
		if item.Nick == data.User.Nick || item.Email == data.User.Email {
			return true
		}
		return false
	})
	if has {
		return nil, errors.New("conflict fields")
	}
	u.db.Users = append(u.db.Users, &data.User)

	if data.Token != nil {
		token := models.Token{Id: data.Token.Id, Value: data.Token.Value}
		u.db.Tokens = append(u.db.Tokens, &token)
	}

	return data, nil
}

func (u *userRepository) GetById(ctx context.Context, id string) (*aggregate.UserAggregate, error) {
	var searched *models.User = nil
	for _, user := range u.db.Users {
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
	lo.Map(u.db.Users, func(user *models.User, i int) *models.User {
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
	lo.Filter(u.db.Users, func(user *models.User, index int) bool {
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
	return lo.ContainsBy(u.db.Users, func(item *models.User) bool {
		if item.Nick == nick {
			return true
		}
		return false
	}), nil
}

func (u *userRepository) HasUserByEmail(ctx context.Context, email string) (bool, error) {
	return lo.ContainsBy(u.db.Users, func(item *models.User) bool {
		if item.Email == email {
			return true
		}
		return false
	}), nil
}

func (u *userRepository) GetByNick(ctx context.Context, nick string) (*aggregate.UserAggregate, error) {
	user, ok := lo.Find(u.db.Users, func(item *models.User) bool {
		if item.Nick == nick {
			return true
		}
		return false
	})
	if ok {
		return &aggregate.UserAggregate{User: *user}, nil
	}
	return nil, nil
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{inMemDb.New()}
}
