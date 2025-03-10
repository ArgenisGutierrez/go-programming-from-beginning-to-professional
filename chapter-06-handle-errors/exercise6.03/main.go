// Package main creating an application to calculate pay for the week
package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hour worked")
)

func main() {
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}
	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime, nil
	}

	return hoursWorked * hourlyRate, nil
}
