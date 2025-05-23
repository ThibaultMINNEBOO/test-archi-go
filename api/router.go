package api

import (
	"net/http"

	"github.com/ThibaultMINNEBOO/apigo/api/users"
)

func RegisterRoutes(mux *http.ServeMux) {
	users.RegisterRoutes(mux, "/api/users")
}
