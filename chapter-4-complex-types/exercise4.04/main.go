// Package main reading a single item from an array
package main

import "fmt"

func message() string {
	arr := [...]string{
		"readey",
		"Get",
		"Go",
		"To",
	}
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Println(message())
}
