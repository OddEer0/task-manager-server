package repository

import (
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
	"github.com/OddEer0/task-manager-server/pkg/shared"
)

type ProjectRepository interface {
	shared.CRUDRepository[*aggregate.ProjectAggregate, dto.CreateProjectDto]
}
