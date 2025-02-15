// Package main creating a custom error message for a banking application
package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid las name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func main() {
	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingNumber)
}
