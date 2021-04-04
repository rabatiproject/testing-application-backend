package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
	"github.com/rabatiproject/testing-application-backend/controllers"
	"github.com/rabatiproject/testing-application-backend/middlewares"
	"log"
	"net/http"
)

func main() {
	log.Println("Application Started")

	r := mux.NewRouter()

	r.HandleFunc("/signIn", controllers.SignIn).Methods("GET")
	r.HandleFunc("/signUp", controllers.AddUser).Methods("POST")
	r.Use(middlewares.LoggingMiddleware)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.AuthorizationMiddleware)
	apiRouter.HandleFunc("/test", controllers.Test)

	apiRouter.HandleFunc("/exam", controllers.CreateExam).Methods(http.MethodPost)
	apiRouter.HandleFunc("/exam/{exam-id}", controllers.GetExam).Methods(http.MethodGet)
	apiRouter.HandleFunc("/exam/{exam-id}/{question-id}", controllers.AttachQuestionToExam).Methods(http.MethodPut)

	apiRouter.HandleFunc("/question/multipleChoice", controllers.CreateMultipleChoiceQuestion).Methods(http.MethodPost)
	apiRouter.HandleFunc("/question/openEnded", controllers.CreateOpenEndedQuestion).Methods(http.MethodPost)
	apiRouter.HandleFunc("/question/programming", controllers.CreateProgrammingQuestion).Methods(http.MethodPost)

	printDynamoTables()

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func printDynamoTables() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	input := &dynamodb.ListTablesInput{}
	fmt.Printf("Tables:\n")
	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		for _, n := range result.TableNames {
			fmt.Println(*n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}

	}
}
