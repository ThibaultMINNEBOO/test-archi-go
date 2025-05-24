package users

import (
	"encoding/json"
	"net/http"

	"github.com/ThibaultMINNEBOO/apigo/api/base"
)

func HandleFindAll(w http.ResponseWriter, r *http.Request) {
	users := make(chan []User)
	go FindAllUsers(users)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(<-users)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDto CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&createUserDto)
	user := make(chan User)

	if err != nil || createUserDto.Name == "" || createUserDto.Age <= 0 {
		httpError := base.HTTPError{Status: 400, Message: "Data invalid"}
		base.JsonError(w, httpError, 400)
		return
	}

	go CreateUser(createUserDto, user)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(<-user)
}
