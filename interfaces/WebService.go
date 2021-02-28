package interfaces

import "net/http"

type HttpHandler func(http.ResponseWriter, *http.Request)

type WebService struct {
	Path    string
	Handler HttpHandler
}

func (service *WebService) Register() {
	http.HandleFunc(service.Path, service.Handler)
}
