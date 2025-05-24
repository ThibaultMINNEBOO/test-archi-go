package users

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, path string) {
	mux.HandleFunc(fmt.Sprintf("GET %s/", path), HandleFindAll)
	mux.HandleFunc(fmt.Sprintf("POST %s/", path), HandleCreateUser)
}
