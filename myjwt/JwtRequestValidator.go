package myjwt

import "github.com/rabatiproject/testing-application-backend/model"

type JwtValidator struct {
}

func (validator *JwtValidator) Validate(credential model.UserCredential) bool {
	if credential.Username == "deniz" && credential.Password == "gursoy" {
		return true
	} else {
		return false
	}
}
