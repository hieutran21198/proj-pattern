package model

// User is the model of user
type User struct {
	Email    string
	Password []byte
	Salt     []byte
}
