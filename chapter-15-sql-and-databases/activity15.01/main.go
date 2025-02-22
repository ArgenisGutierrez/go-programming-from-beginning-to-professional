// Package main holding user data in a table
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

var (
	userTable = `
  CREATE TABLE User (
    id integer not null,
    name varchar(30),
    email varchar(30)
  )
  `
	createUser = `
  INSERT INTO User VALUES (?,?,?)
  `
	updateUser = `
  UPDATE User
  SET email = ?
  WHERE id=?
  `
	deleteUser = `
  DELETE FROM User
  WHERE id=?
  `
)

func main() {
	user1 := User{ID: 1, Name: "Arfhel", Email: "arfhel@gmail.com"}
	user2 := User{ID: 2, Name: "Samael", Email: "samael@gmail.com"}
	users := []User{user1, user2}

	db, err := sql.Open("mysql", "arfhel:sheyla95@/go")
	if err != nil {
		panic(err)
	}

	// Create the table User
	_, err = db.Exec(userTable)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table created successfully")
	}

	insert, err := db.Prepare(createUser)
	if err != nil {
		panic(err)
	}
	for _, v := range users {
		_, err = insert.Exec(v.ID, v.Name, v.Name)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("The user with name: %s and email: %s was successfully added!", v.Name, v.Email)
		}
	}

	update, err := db.Prepare(updateUser)
	if err != nil {
		panic(err)
	}
	_, err = update.Exec("user@packt.com", 1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The user email address was successfully updated")
	}

	delete, err := db.Prepare(deleteUser)
	if err != nil {
		panic(err)
	}
	_, err = delete.Exec(2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The second user was successfully removed!")
	}

	defer db.Close()

}
