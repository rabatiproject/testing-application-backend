package definitions

import "github.com/rabatiproject/testing-application-backend/model/base"

type UserRepo interface {
	CreateUser(user *base.User) error
	UserExists(email string) bool
	GetUserFrom(email string) *base.User
}
