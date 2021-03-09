package main

import (
	"github.com/gorilla/mux"
	"github.com/rabatiproject/testing-application-backend/middlewares"
	"github.com/rabatiproject/testing-application-backend/services"
	"log"
	"net/http"
)

func main() {
	log.Println("Application Started")

	r := mux.NewRouter()
	r.HandleFunc("/SignIn", services.SignIn).Methods("GET")
	r.Use(middlewares.LoggingMiddleware)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
