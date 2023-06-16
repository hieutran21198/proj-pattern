package dto

type (
	// LoginRequest is the input data for the business logic of Login.
	LoginRequest struct {
		Username *string `json:"username" validate:"required,min=6"`
		Password *string `json:"password" validate:"required,min=6"`
	}

	// LoginResponse is the output data for the business logic of Login.
	LoginResponse struct {
		Success bool `json:"success"`
	}
)
