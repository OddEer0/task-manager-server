package aggregate

import "task-manager-server/internal/domain/models"

type UserAggregate struct {
	User     models.User
	Token    *models.Token
	Projects []*models.Project
}

func (u *UserAggregate) SetToken(token *models.Token) {
	u.Token = token
}
