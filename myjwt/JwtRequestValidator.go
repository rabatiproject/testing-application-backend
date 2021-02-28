package myjwt

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

//const token_bearer = "Bearer"

type JwtRequestValidator struct {
}

func (jwtValidator JwtRequestValidator) IsValid(req *http.Request) bool {
	authHeader := req.Header.Get("Authorization")

	if authHeader == "" {
		return false
	} else {
		if len(authHeader) > 0 {
			//	tknStr := strings.ReplaceAll(authHeader, token_bearer, "")

			claims := &jwt.StandardClaims{}

			tkn, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					return false
				}
				return false
			}
			if !tkn.Valid {
				return false
			}

		} else {
			return false
		}
	}
	return true
}
