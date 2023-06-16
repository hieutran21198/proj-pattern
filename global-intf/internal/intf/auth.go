package intf

import (
	"context"

	"global-intf/internal/dto"
)

// AuthService represents the interface of auth service.
type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (res *dto.LoginResponse, err error)
}

// JWTStrategy represents the interface of jwt strategy.
type JWTStrategy interface {
}

// MFAStrategy represents the interface of mfa strategy.
type MFAStrategy interface {
}
