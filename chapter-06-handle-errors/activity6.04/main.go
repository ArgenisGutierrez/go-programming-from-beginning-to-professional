// Package main preventing a panic form crashing the app
package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d directDeposit) validateRountingNumber() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
	return nil
}
func (d directDeposit) validateLastName() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	lastName := strings.TrimSpace(d.lastName)
	if len(lastName) == 0 {
		panic(ErrInvalidLastName)
	}
	return nil
}

func (d directDeposit) report() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Printf("Last Name: %s\n", d.lastName)
	fmt.Printf("First Name: %s\n", d.firstName)
	fmt.Printf("Bank Name: XYZ Inc\n")
	fmt.Printf("Rounting Number: %d\n", d.routingNumber)
	fmt.Printf("Accounting Number : %d\n", d.accountNumber)
}

func main() {
	dd := directDeposit{
		lastName:      "  ",
		firstName:     "Abe",
		bankName:      "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := dd.validateRountingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	dd.report()

}
