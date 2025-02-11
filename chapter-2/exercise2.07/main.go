// Package main expressionless switch statements
package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born in weekend")
	default:
		fmt.Println("Born some other day")
	}
}
