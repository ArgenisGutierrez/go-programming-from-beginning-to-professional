// Package main Numeric Operators
package main

import "fmt"

func main() {
	// First 2 Item
	var total float64 = 2 * 13
	fmt.Println("Sub: ", total)
	// 2 Items more
	total += 4 * 2.25
	// Have a discount
	total -= 5
	fmt.Println("Sub: ", total)
	// Add the tip
	tip := total * 0.1
	total += tip
	fmt.Println("Total: ", total)
	// Split the bill
	split := total / 2
	fmt.Println("Split: ", split)
	//Reward every 5th visit
	visitCount := 24
	visitCount += 1

	//Verfify if is the visit No 5
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward'")
	}
}
