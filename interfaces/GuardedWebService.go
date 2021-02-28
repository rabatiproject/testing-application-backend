package interfaces

import (
	"net/http"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

type GuardedService struct {
	Path    string
	handler HttpHandler
}

func (service *GuardedService) register() {
	http.HandleFunc(service.Path, service.handler)
}
func wrapp(){
	
}
