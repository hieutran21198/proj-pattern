package crypto

import "global-intf/internal/intf"

// Service represents the crypto service.
type Service struct {
}

// to check the 'Service' is missing methods from the 'intf.CryptoService' or not.
var _ intf.CryptoService = &Service{}

// NewService creates a Service.
func NewService() *Service {
	return &Service{}
}

// ComparePassword implements intf.CryptoService.
func (*Service) ComparePassword([]byte, []byte, []byte) error {
	panic("unimplemented")
}

// GenUlid implements intf.CryptoService.
func (*Service) GenUlid() string {
	panic("unimplemented")
}

// HashPassword implements intf.CryptoService.
func (*Service) HashPassword(password []byte) ([]byte, string, error) {
	panic("unimplemented")
}
