// Package main finding the messages of specific users
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	messagesTable = `
  CREATE TABLE Message (
    UserID integer(10),
    Message varchar(280)
  )
  `
	messagesInsert = `
  INSERT INTO Message VALUES (?,?)
  `
	messageSelect = `
  SELECT User.name, User.email, Message.Message 
  FROM Message 
  JOIN User
  ON User.id=Message.UserID
  WHERE User.name LIKE ?
  `
)

type Message struct {
	UserID  int
	Message string
}

func main() {
	messages := []Message{
		{UserID: 1, Message: "Message of proof 1"},
		{UserID: 1, Message: "Message of proof 2"},
		{UserID: 3, Message: "Message of proof 3"},
		{UserID: 5, Message: "Message of proof 4"},
	}
	db, err := sql.Open("mysql", "arfhel:sheyla95@/go")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection successfully")
	}

	_, err = db.Exec(messagesTable)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table was created successfully")
	}

	insert, err := db.Prepare(messagesInsert)
	if err != nil {
		panic(err)
	}
	for _, v := range messages {
		_, err = insert.Exec(v.UserID, v.Message)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("The UsersID: %d with message: %s was successfully added!\n", v.UserID, v.Message)
		}
	}

	sMessage, err := db.Prepare(messageSelect)
	if err != nil {
		panic(err)
	}

  var (
    userName string
    userEmail string
    userMessage string
  )
  result ,err := sMessage.Query("arfhel")
  for result.Next(){
    err = result.Scan(&userName,&userEmail,&userMessage)
    if err !=nil{
      panic(err)
    }
    fmt.Printf("The user: %s with email: %s has following message: %s\n",userName,userEmail,userMessage)
  }

  defer db.Close()
}
