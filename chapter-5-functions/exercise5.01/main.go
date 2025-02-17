// Package main creating a function print salesperson expectation ratings from the number of items sold
package main

import "fmt"

func main() {
	itemSold()
}

func itemSold() {
	items := make(map[string]int)
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println("is below expectation")
		} else if v > 40 && v < 100 {
			fmt.Println("meets expectation")
		} else if v > 100 {
			fmt.Println("Exceeded expectation")
		}
	}
}
