package inMemDb

import "github.com/OddEer0/task-manager-server/internal/domain/models"

type InMemDb struct {
	Users    []*models.User
	Tokens   []*models.Token
	Projects []*models.Project
	Columns  []*models.Column
	Tasks    []*models.Task
	Tags     []*models.Tag
}

var instance *InMemDb = nil

func New() *InMemDb {
	if instance != nil {
		return instance
	}

	instance = &InMemDb{
		Users:    []*models.User{},
		Tokens:   []*models.Token{},
		Projects: []*models.Project{},
		Columns:  []*models.Column{},
		Tasks:    []*models.Task{},
		Tags:     []*models.Tag{},
	}

	return instance
}
