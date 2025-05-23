package users

import (
	"fmt"

	"github.com/ThibaultMINNEBOO/apigo/api/database"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func FindAllUsers(ch chan []User) {
	db := database.GetDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM users")

	if err != nil {
		fmt.Println(err)
		return
	}

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			fmt.Printf("Error during scan : %s", err)
			return
		}
		users = append(users, u)
	}

	ch <- users
}
