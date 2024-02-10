package mock_repository

import (
	"context"
	"errors"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/models"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
	inMemDb "github.com/OddEer0/task-manager-server/internal/infrastructure/storage/in_mem_db"
	"github.com/samber/lo"
)

type projectRepository struct {
	db *inMemDb.InMemDb
}

func (p projectRepository) Create(ctx context.Context, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error) {
	has := lo.ContainsBy(p.db.Projects, func(item *models.Project) bool {
		if item.Name == data.Project.Name {
			return true
		}
		return false
	})
	if has {
		return nil, errors.New("conflict fields")
	}
	p.db.Projects = append(p.db.Projects, &data.Project)
	return data, nil
}

func (p projectRepository) GetById(ctx context.Context, id string) (*aggregate.ProjectAggregate, error) {
	res, ok := lo.Find(p.db.Projects, func(project *models.Project) bool {
		if project.Id == id {
			return true
		}
		return false
	})
	if !ok {
		return nil, errors.New("not found")
	}
	return &aggregate.ProjectAggregate{Project: *res}, nil
}

func (p projectRepository) Update(ctx context.Context, id string, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error) {
	has := lo.ContainsBy(p.db.Projects, func(item *models.Project) bool {
		if item.Name == data.Project.Name {
			return true
		}
		return false
	})
	if !has {
		return nil, errors.New("not found")
	}

	lo.Map(p.db.Projects, func(project *models.Project, index int) *models.Project {
		if id == project.Id {
			return &data.Project
		}
		return project
	})

	return data, nil
}

func (p projectRepository) Delete(ctx context.Context, id string) error {
	has := false
	lo.Filter(p.db.Projects, func(project *models.Project, index int) bool {
		if project.Id != id {
			return true
		}
		has = true
		return false
	})

	if has {
		return nil
	} else {
		return errors.New("not found")
	}
}

func NewProjectRepository() repository.ProjectRepository {
	return &projectRepository{
		db: inMemDb.New(),
	}
}
