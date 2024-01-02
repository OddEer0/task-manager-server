package repository

import (
	"github.com/OddEer0/task-manager-server/internal/domain/aggregate"
	"github.com/OddEer0/task-manager-server/internal/pkg/shared"
	"github.com/OddEer0/task-manager-server/internal/presentation/dto"
)

type TagRepository struct {
	shared.CRUDRepository[*aggregate.TagAggregate, dto.CreateTagDto]
}
