// Package main Use of GORM go ORM
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

func main() {
	dsn := "arfhel:sheyla95@tcp/go"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection successfully")
	}
	db.AutoMigrate(&User{})

	u := &User{FirstName: "John", LastName: "Smith", Email: "john.smith@gmail.com"}
	db.Create(&u)

	db.Create(&User{FirstName: "James", LastName: "Doe", Email: "james.doe@gmail.com"})

	var user User
	db.First(&user, "last_name=?", "Doe")
}
