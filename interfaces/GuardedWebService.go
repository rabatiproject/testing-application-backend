package interfaces

import (
	"github.com/rabatiproject/testing-application-backend/jwtConfigs"
	"net/http"
)

type GuardedWebService struct {
	WebService
	Validator RequestValidator
}

var (
	jwtValidator jwtConfigs.JwtRequestValidator
)

func NewGuardedService(webService *WebService) *GuardedWebService {
	guardedWebService := &GuardedWebService{}

	guardedWebService.Path = webService.Path
	guardedWebService.Handler = func(writer http.ResponseWriter, req *http.Request) {
		if jwtValidator.IsValid(req) {
			webService.Handler(writer, req)
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	guardedWebService.Validator = jwtValidator
	return guardedWebService
}
