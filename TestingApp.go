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

	r.HandleFunc("/signIn", services.SignIn).Methods("GET")
	r.Use(middlewares.LoggingMiddleware)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.AuthorizationMiddleware)
	apiRouter.HandleFunc("/test", services.Test)

	apiRouter.HandleFunc("/exam", services.CreateExam).Methods(http.MethodPost)
	apiRouter.HandleFunc("/exam/{id}", services.DeleteExam).Methods(http.MethodDelete)

	apiRouter.HandleFunc("/question/multipleChoice", services.CreateMultipleChoiceQuestion).Methods(http.MethodPost)
	apiRouter.HandleFunc("/question/openEnded", services.CreateOpenEndedQuestion).Methods(http.MethodPost)
	apiRouter.HandleFunc("/question/programming", services.CreateProgrammingQuestion).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
