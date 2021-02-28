package services

import (
	"fmt"
	"net/http"
)

func Test(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("Has reached")

}
