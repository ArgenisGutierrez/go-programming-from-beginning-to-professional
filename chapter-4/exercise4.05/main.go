// Package main writing to an array
package main

import "fmt"

func message() string {
	arr := [4]string{"ready", "Get", "Go", "to"}
	arr[0] = "Its"
	arr[0] = "time"
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}
func main() {

	fmt.Println(message())
}
