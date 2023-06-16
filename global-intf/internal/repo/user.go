package repo

import (
	"context"

	"gorm.io/gorm"

	"global-intf/internal/intf"
	"global-intf/internal/model"
)

// UserRepo is the implementation of UserRepo.
type UserRepo struct {
	db *gorm.DB

	// implementation crud
	CRUD[model.User]
}

var _ intf.UserRepository = &UserRepo{}

var _ intf.CRUD[model.User] = &UserRepo{}

func (u *UserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
