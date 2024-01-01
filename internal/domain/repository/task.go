package repository

import (
	"task-manager-server/internal/domain/aggregate"
	"task-manager-server/internal/pkg/shared"
	"task-manager-server/internal/presentation/dto"
)

type TaskRepository interface {
	shared.CRUDRepository[*aggregate.TaskAggregate, dto.CreateTaskDto]
}
