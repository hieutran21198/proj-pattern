package auth

import (
	"global-intf/internal/intf"
)

// Service represents the  service.
type Service struct {
	userRepo intf.UserRepo

	crypto intf.CryptoService
}

// to check the 'Service' is missing methods from the 'intf.AuthService' or not.
var _ intf.AuthService = &Service{}

// NewService creates a Service.
func NewService(userRepo intf.UserRepo, crypto intf.CryptoService) *Service {
	return &Service{
		userRepo: userRepo,
		crypto:   crypto,
	}
}
