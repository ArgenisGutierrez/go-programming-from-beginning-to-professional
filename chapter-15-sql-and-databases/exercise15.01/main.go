// Package main creating a table that hold a series of numbers
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// create the db connection
	db, err := sql.Open("mysql", "arfhel:sheyla95@/go")
	// check if the connection was successfully
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection was successfully")
	}
	defer db.Close()

	// query to create the table
	TableCreate := `
  CREATE TABLE Numbers (
    Number integer(5),
    Property varchar(10)
  )
  `
	// execute the query to create the table
	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table Numbers was successfully created")
	}

	// Prepare prevent SQL injections wrapped our queries
	insert, err := db.Prepare("INSERT INTO Numbers VALUES (?,?)")
	if err != nil {
		panic(err)
	}

	var prop string

	// check if a number is oddd or even
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "Even"
		} else {
			prop = "Odd"
		}

		// execute the query to insert data
		_, err = insert.Exec(i, prop)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("The number: ", i, "is:", prop)
		}
	}
	defer insert.Close()
	fmt.Println("The numbers are ready")

}
