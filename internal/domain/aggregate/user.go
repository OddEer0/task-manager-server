package aggregate

import "github.com/OddEer0/task-manager-server/internal/domain/models"

type UserAggregate struct {
	User     models.User
	Token    *models.Token
	Projects []*models.Project
}

func (u *UserAggregate) Validation() error {
	if err := u.User.Birthday.Validate(); err != nil {
		return err
	}
	if err := u.User.Email.Validate(); err != nil {
		return err
	}
	return nil
}

func (u *UserAggregate) SetToken(token *models.Token) {
	u.Token = token
}

func NewUserAggregate(user models.User) (*UserAggregate, error) {
	result := UserAggregate{User: user}
	if err := result.Validation(); err != nil {
		return nil, err
	}
	return &result, nil
}
