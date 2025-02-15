// Package main Crashing the program on errors using a panic
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
  pay := payDay(81,50)
  fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int {
  report := func ()  {
    fmt.Printf("hoursWorked : %d\nHourlyRat : %d\n",hoursWorked,hourlyRate)
  }
  defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
    panic(ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
    panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}

	return hoursWorked * hourlyRate
}
