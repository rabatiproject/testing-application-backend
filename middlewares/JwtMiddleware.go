package middlewares

import (
	"log"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Checking authorization")
		next.ServeHTTP(writer, request)
	})
}
