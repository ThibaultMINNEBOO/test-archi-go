package users

import (
	"encoding/json"
	"net/http"
)

func HandleFindAll(w http.ResponseWriter, r *http.Request) {
	users := make(chan []User)
	go FindAllUsers(users)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(<-users)
}
