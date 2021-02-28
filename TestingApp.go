package main

import (
	"github.com/rabatiproject/testing-application-backend/interfaces"
	"github.com/rabatiproject/testing-application-backend/myjwt"
	"github.com/rabatiproject/testing-application-backend/services"
	"log"
	"net/http"
)

func main() {
	log.Println("Application Started")

	jwtValidator := myjwt.JwtRequestValidator{}

	signInService := interfaces.WebService{Path: "/SignIn", Handler: services.SignIn}
	signInService.Register()

	testService := interfaces.NewGuardedService(&interfaces.WebService{
		Path:    "/Test",
		Handler: services.Test,
	}, jwtValidator)
	testService.Register()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Server could not be initiated")
	}

}
