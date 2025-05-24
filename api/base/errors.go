package base

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func JsonError(w http.ResponseWriter, err HTTPError, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
