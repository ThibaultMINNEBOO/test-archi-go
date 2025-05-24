package users

import (
	"fmt"

	"github.com/ThibaultMINNEBOO/apigo/api/database"
)

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

func CreateUser(createUser CreateUserDTO, userChannel chan User) {
	db := database.GetDatabase()
	defer db.Close()

	tx, _ := db.Begin()
	stmt, err := tx.Prepare("INSERT INTO users (name, age) VALUES (?, ?)")

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := stmt.Exec(createUser.Name, createUser.Age)

	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	lastInsertedId, _ := res.LastInsertId()

	tx.Commit()

	user := User{Id: int(lastInsertedId), Name: createUser.Name, Age: createUser.Age}

	userChannel <- user
}
