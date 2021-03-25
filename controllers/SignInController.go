package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rabatiproject/testing-application-backend/beans"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"github.com/rabatiproject/testing-application-backend/model/jwtModel"
	"github.com/rabatiproject/testing-application-backend/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	JwtSecretKey     = []byte("my_secret_key")
	AuthHeaderPrefix = "Bearer "
	AuthHeader       = "Authorization"
)

type JwtResponse struct {
	User *base.User `json:"user"`
	jwt.StandardClaims
}

func SignIn(writer http.ResponseWriter, r *http.Request) {

	var credential jwtModel.UserCredential
	parseError := json.NewDecoder(r.Body).Decode(&credential)

	if parseError != nil &&
		len(strings.TrimSpace(credential.Email)) == 0 &&
		utils.IsEmailValid(credential.Email) &&
		len(strings.TrimSpace(credential.Password)) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println(parseError.Error())
	} else {
		user := beans.UserRepository.GetUserFrom(credential.Email)
		if credential.Email == user.Email && credential.Password == user.Password {

			expirationTime := time.Now().Add(5 * time.Hour)
			jwtResponse := &JwtResponse{
				User: user,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtResponse)
			tokenString, error := token.SignedString(JwtSecretKey)

			if error == nil {
				log.Println(user.Id + " signed in successfully")
				writer.Header().Add(AuthHeader, AuthHeaderPrefix+tokenString)
				return
			}

		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

	}

}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	user := &base.User{}
	json.NewDecoder(request.Body).Decode(user)

	if len(strings.TrimSpace(user.Surname)) == 0 &&
		len(strings.TrimSpace(user.Name)) == 0 &&
		len(strings.TrimSpace(user.Email)) == 0 &&
		utils.IsEmailValid(user.Email) &&
		beans.UserRepository.UserExists(user.Email) {

		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	creationError := beans.UserRepository.CreateUser(user)

	if creationError != nil {
		writer.WriteHeader(http.StatusNoContent)
		return
	}
}
