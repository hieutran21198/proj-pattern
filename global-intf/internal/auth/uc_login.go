package auth

import (
	"context"

	"gorm.io/gorm"

	"global-intf/internal/dto"
	"global-intf/pkg/server"
)

// Login implements intf.LoginService
func (s *Service) Login(ctx context.Context, req *dto.LoginRequest) (res *dto.LoginResponse, err error) {
	res = &dto.LoginResponse{}

	user, err := s.userRepo.FindByEmail(ctx, *req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrInvalidCredential
		}

		return nil, server.ErrInternalDatabaseError.WithInternal(err)
	}

	if err = s.crypto.ComparePassword(user.Password, user.Salt, []byte(*req.Password)); err != nil {
		return nil, ErrInvalidCredential
	}

	res.Success = true
	return res, nil
}
