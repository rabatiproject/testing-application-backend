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

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.AuthorizationMiddleware)
	apiRouter.HandleFunc("/Test", services.Test)

	apiRouter.HandleFunc("/Exam", services.CreateExam).Methods(http.MethodPost)
	apiRouter.HandleFunc("/Exam/{id}", services.DeleteExam).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
