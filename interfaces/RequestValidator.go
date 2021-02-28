package interfaces

import "net/http"

type RequestValidator interface {
	IsValid(*http.Request) bool
}
