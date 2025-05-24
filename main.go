package main

import (
	"log"
	"net/http"

	"github.com/ThibaultMINNEBOO/apigo/api"
	"github.com/ThibaultMINNEBOO/apigo/api/database"
)

func main() {
	mux := http.NewServeMux()

	db := database.GetDatabase()
	defer db.Close()
	database.InitDatabase(db)
	api.RegisterRoutes(mux)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}
