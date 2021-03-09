package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rabatiproject/testing-application-backend/services"

	"net/http"
	"strings"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		token := request.Header.Get(services.AuthHeader)
		token = strings.ReplaceAll(token, services.AuthHeaderPrefix, "")

		if len(strings.TrimSpace(token)) == 0 {
			writer.WriteHeader(http.StatusBadRequest)
			return
		} else {

			tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error.")
				}
				return services.JwtSecretKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					writer.WriteHeader(http.StatusBadRequest)
					return
				}
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			if !tkn.Valid {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}

		}
		next.ServeHTTP(writer, request)
	})
}
