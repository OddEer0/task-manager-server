package repository

import (
	"task-manager-server/internal/domain/aggregate"
	"task-manager-server/internal/pkg/shared"
	"task-manager-server/internal/presentation/dto"
)

type TagRepository struct {
	shared.CRUDRepository[*aggregate.TagAggregate, dto.CreateTagDto]
}
