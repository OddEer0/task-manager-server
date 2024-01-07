package shared

import "context"

type CRUDRepository[A any, C any] interface {
	Create(ctx context.Context, data C) (A, error)
	GetById(ctx context.Context, id string) (A, error)
	Update(ctx context.Context, id string, data A) (A, error)
	Delete(ctx context.Context, id string) error
}
