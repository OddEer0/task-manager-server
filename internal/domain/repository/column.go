package repository

import (
	"task-manager-server/internal/domain/aggregate"
	"task-manager-server/internal/pkg/shared"
	"task-manager-server/internal/presentation/dto"
)

type ColumnRepository interface {
	shared.CRUDRepository[*aggregate.ColumnAggregate, dto.CreateColumnDto]
}
