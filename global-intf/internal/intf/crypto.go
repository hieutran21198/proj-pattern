package intf

// CryptoService represents the interface of crypto service.
type CryptoService interface {
	// HashPassword returns hashed password with salt.
	HashPassword(password []byte) ([]byte, string, error)

	// ComparePassword returns true if the password is correct.
	ComparePassword(hashed []byte, salt []byte, password []byte) error

	// GenUlid generates an ulid that must be safe for concurrent use.
	GenUlid() string
}
