package aggregate

import (
	"github.com/OddEer0/task-manager-server/internal/common/lib/app_validator"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
)

type UserAggregate struct {
	User     models.User
	Token    *models.Token
	Projects []*models.Project
}

func (u *UserAggregate) SetToken(refreshToken string) error {
	token := models.Token{Id: u.User.Id, Value: refreshToken}
	validator := appValidator.New()
	err := validator.Struct(token)
	if err != nil {
		return err
	}
	u.Token = &token
	return nil
}

func (u *UserAggregate) Validation() error {
	validator := appValidator.New()
	err := validator.Struct(u.User)
	if err != nil {
		return err
	}
	return nil
}

func NewUserAggregate(user models.User) (*UserAggregate, error) {
	result := UserAggregate{User: user}
	if err := result.Validation(); err != nil {
		return nil, err
	}
	return &result, nil
}
