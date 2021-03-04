package myjwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/rabatiproject/testing-application-backend/model"
	"net/http"
	"time"
)

var (
	JwtSecretKey = []byte("my_secret_key")
)

type JwtResponse struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(credential model.UserCredential) (string, int) {

	expirationTime := time.Now().Add(5 * time.Hour)
	jwtResponse := &JwtResponse{
		Username: credential.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtResponse)
	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		return "", http.StatusInternalServerError
	} else {
		return tokenString, http.StatusOK
	}
}
