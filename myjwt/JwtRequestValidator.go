package myjwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type JwtRequestValidator struct {
}

func (jwtValidator JwtRequestValidator) IsValid(req *http.Request) bool {
	authHeader := req.Header.Get("Authorization")

	if authHeader == "" {
		return false
	} else {
		if len(authHeader) > 0 {
			//	tknStr := strings.ReplaceAll(authHeader, token_bearer, "")

			//claims := &jwt.StandardClaims{}

			tkn, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return JwtSecretKey, nil
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
