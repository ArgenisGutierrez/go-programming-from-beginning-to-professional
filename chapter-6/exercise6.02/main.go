// Package main a semantic error with walking distance
package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrBadData := errors.New("some bad data")
	fmt.Println(ErrBadData)
	km := 2
	if km > 2 {
		fmt.Println("Take the car")
	} else {
		fmt.Println("Going to walk today")
	}
}
