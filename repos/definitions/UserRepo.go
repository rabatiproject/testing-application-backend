package definitions

import "github.com/rabatiproject/testing-application-backend/model/base"

type UserRepo interface {
	CreateUser(user *base.User) error
}
