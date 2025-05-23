package users

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, path string) {
	mux.HandleFunc(fmt.Sprintf("%s/", path), HandleFindAll)
}
