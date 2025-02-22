// Package main holding prime numbers in a database
package main

import (
	"database/sql"
	"fmt"
	"math/big"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		number    int64
		prop      string
		primeSum  int64
		newNumber int64
	)
	db, err := sql.Open("mysql", "user:password@/db")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection was successfully")
	}

	allNumbers := "SELECT * FROM Numbers"
	numbers, err := db.Prepare(allNumbers)
	if err != nil {
		panic(err)
	}

	primeSum = 0
	result, err := numbers.Query()
	fmt.Println("The list of prime numbers:")
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		if big.NewInt(number).ProbablyPrime(0) {
			primeSum += number
			fmt.Println(" ", number)
		}
	}
	err = numbers.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nThe total sum of prime numbers in this range is:", primeSum)

	remove := "DELETE FROM Numbers WHERE Property=?"
	removeResult, err := db.Exec(remove, "Even")
	if err != nil {
		panic(err)
	}
	modifiedRecords, err := removeResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("The number of rows removed:", modifiedRecords)
	fmt.Println("Updating numbers...")
	update := "UPDATE Numbers SET Number=? WHERE Number=? AND Property=?"
	allTheNumbers := "SELECT * FROM Numbers"
	numbers, err = db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	result, err = numbers.Query()
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		newNumber = number + primeSum
		_, err = db.Exec(update, newNumber, number, prop)
		if err != nil {
			panic(err)
		}
	}
	numbers.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("The execution is now complete...")
	db.Close()
}
