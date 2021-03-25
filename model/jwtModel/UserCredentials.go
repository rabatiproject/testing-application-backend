package jwtModel

type UserCredential struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
