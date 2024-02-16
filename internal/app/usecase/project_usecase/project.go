package projectUsecase

import (
	"context"
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/domain/repository"
)

type (
	ProjectUseCase interface {
		Create(ctx context.Context, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error)
		GetById(ctx context.Context, id string) (*aggregate.ProjectAggregate, error)
		Update(ctx context.Context, id string, data *aggregate.ProjectAggregate) (*aggregate.ProjectAggregate, error)
		Delete(ctx context.Context, id string) error
	}

	projectUseCase struct {
		repository.ProjectRepository
	}
)

func New(projectRepository repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{ProjectRepository: projectRepository}
}
