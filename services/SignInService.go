package services

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rabatiproject/testing-application-backend/model/jwtModel"
	"log"
	"net/http"
	"time"
)

var (
	JwtSecretKey     = []byte("my_secret_key")
	AuthHeaderPrefix = "Bearer "
	AuthHeader       = "Authorization"
)

type JwtResponse struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignIn(writer http.ResponseWriter, r *http.Request) {

	var credential jwtModel.UserCredential
	parseError := json.NewDecoder(r.Body).Decode(&credential)

	if parseError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Not valid"))
		fmt.Println(parseError.Error())
	} else {

		if credential.Username == "deniz" && credential.Password == "deniz" {

			expirationTime := time.Now().Add(5 * time.Hour)
			jwtResponse := &JwtResponse{
				Username: credential.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtResponse)
			tokenString, error := token.SignedString(JwtSecretKey)

			if error == nil {
				log.Println(credential.Username + " signed in successfully")
				writer.Header().Add(AuthHeader, AuthHeaderPrefix+tokenString)
				return
			}

		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

	}

}
