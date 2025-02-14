// Package main creating a closure function to decrement a counter
package main

import "fmt"

var counter = 4

func main() {
	x := decrement(counter)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

}

func decrement(i int) func() int {
	return func() int {
		i--
		return i
	}
}
