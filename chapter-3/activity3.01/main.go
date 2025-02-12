// Package main Sales tax calculator
package main

import "fmt"

func main() {
	var product float64 = 0.99
	switch product {
	case 0.99:
		fmt.Println("Sales Tax Total:", product*1.075)
	case 2.75:
		fmt.Println("Sales Tax Total:", product*1.015)
	case 0.87:
		fmt.Println("Sales Tax Total:", product*1.02)
	default:
		fmt.Println("Product not found")
	}
}
