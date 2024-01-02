package repository

import (
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/pkg/shared"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

type ProjectRepository interface {
	shared.CRUDRepository[*aggregate.ProjectAggregate, dto.CreateProjectDto]
}
