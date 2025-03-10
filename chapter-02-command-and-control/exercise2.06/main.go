// Package main switch statement and multiple case values
package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekedn")
	default:
		fmt.Println("Error, day born not valid")
	}
}
