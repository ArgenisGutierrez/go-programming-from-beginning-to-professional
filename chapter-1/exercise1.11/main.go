// Package main Comparing values
package main

import "fmt"

func main() {
	visits := 15
	fmt.Println("First visit :", visits == 1)
	fmt.Println("Return visit :", visits != 1)
	fmt.Println("Silver memeber", visits > 10 && visits < 21)
	fmt.Println("Golder memeber", visits > 20 && visits < 30)
	fmt.Println("Platinum memeber", visits > 30)
}
