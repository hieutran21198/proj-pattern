package intf

import "context"

// CRUD represents the interface of CRUD repo.
type CRUD[Entity any] interface {
	Create(ctx context.Context, entity *Entity) (*Entity, error)
	Update(ctx context.Context, entity *Entity) (*Entity, error)
	GetByID(ctx context.Context, id string) (*Entity, error)
	HardDelete(ctx context.Context, entity *Entity) error
	Delete(ctx context.Context, entity *Entity) error
}
