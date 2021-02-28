package interfaces

import (
	"net/http"
)

type GuardedWebService struct {
	WebService
	Validator RequestValidator
}

//func NewGuardedWebService(webService WebService) *GuardedWebService {
//	return GuardedWebService{}
//}

func NewGuardedService(webService *WebService, validator RequestValidator) *GuardedWebService {
	guardedWebService := &GuardedWebService{}

	guardedWebService.Path = webService.Path
	guardedWebService.Handler = func(writer http.ResponseWriter, req *http.Request) {
		if validator.IsValid(req) {
			webService.Handler(writer, req)
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	guardedWebService.Validator = validator
	return guardedWebService
}
