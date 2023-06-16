package intf

import (
	"context"

	"global-intf/internal/model"
)

// UserRepo represents the interface of user repository.
type UserRepo interface {
	// FindByEmail finds user by email address.
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
