package services

import (
	"encoding/json"
	"fmt"
	"github.com/rabatiproject/testing-application-backend/jwtConfigs"
	"github.com/rabatiproject/testing-application-backend/model/jwt"
	"log"
	"net/http"
)

func SignIn(writer http.ResponseWriter, r *http.Request) {

	var credential jwt.UserCredential
	parseError := json.NewDecoder(r.Body).Decode(&credential)

	if parseError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Not valid"))
		fmt.Println(parseError.Error())
		//log.Panic("Invalid data")
	} else {

		if credential.Username == "deniz" && credential.Password == "deniz" {

			token, responseCode := jwtConfigs.CreateToken(credential)

			if responseCode != http.StatusInternalServerError {
				log.Println(credential.Username + " signed in successfully")
				//writer.Header().Add("Authorization", "Bearer "+token)
				writer.Header().Add("Authorization", token)
				return
			}

		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

	}

}
