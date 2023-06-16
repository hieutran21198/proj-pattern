package repo

import (
	"context"

	"gorm.io/gorm"

	"global-intf/internal/intf"
)

// CRUD is the implementation of intf.CRUD.
type CRUD[Entity any] struct {
	db *gorm.DB
}

var _ intf.CRUD[any] = &CRUD[any]{}

// NewCRUD creates a new CRUD instance.
func NewCRUD[Entity any](db *gorm.DB) *CRUD[Entity] {
	return &CRUD[Entity]{db: db}
}

// Create implements intf.CRUD.
func (s *CRUD[Entity]) Create(ctx context.Context, entity *Entity) (*Entity, error) {
	if err := s.db.Create(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

// Delete implements intf.CRUD.
func (s *CRUD[Entity]) Delete(ctx context.Context, entity *Entity) error {
	if err := s.db.Delete(entity).Error; err != nil {
		return err
	}

	return nil
}

// GetByID implements intf.CRUD.
func (s *CRUD[Entity]) GetByID(ctx context.Context, id string) (*Entity, error) {
	entity := new(Entity)
	if err := s.db.First(entity, id).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

// HardDelete implements intf.CRUD.
func (s *CRUD[Entity]) HardDelete(ctx context.Context, entity *Entity) error {
	if err := s.db.Unscoped().Delete(entity).Error; err != nil {
		return err
	}

	return nil
}

// Update implements intf.CRUD.
func (s *CRUD[Entity]) Update(ctx context.Context, entity *Entity) (*Entity, error) {
	if err := s.db.Save(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}
