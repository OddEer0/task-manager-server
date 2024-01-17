package userAggregateFactory

import (
	"time"

	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/valuesobject"
)

type (
	UserAggregateFactory    struct{}
	CreateUserAggregateData struct {
		Id             string
		Nick           string
		Password       string
		Email          string
		FirstName      string
		LastName       string
		SubTitle       *string
		Avatar         *string
		Birthday       time.Time
		Role           string
		ActivationLink string
		IsActivate     bool
		IsBanned       bool
		BanReason      *string
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)

func (u UserAggregateFactory) CreateUserAggregate(data CreateUserAggregateData) (*aggregate.UserAggregate, error) {
	hashPassword, err := valuesobject.NewPassword(data.Password)
	if err != nil {
		return nil, err
	}
	var birthday *valuesobject.Birthday = nil
	if !data.Birthday.IsZero() {
		tmp, err := valuesobject.NewBirthday(data.Birthday)
		if err != nil {
			return nil, err
		}
		birthday = &tmp
	}

	return aggregate.NewUserAggregate(models.User{
		Id:             data.Id,
		Nick:           data.Nick,
		Email:          data.Email,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		SubTitle:       data.SubTitle,
		Avatar:         data.Avatar,
		Password:       hashPassword,
		Birthday:       birthday,
		Role:           data.Role,
		ActivationLink: data.ActivationLink,
		IsActivate:     data.IsActivate,
		IsBanned:       data.IsBanned,
		BanReason:      data.BanReason,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	})
}
