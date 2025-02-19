// Package main printing decimal, binary, and hex values
package main

import "fmt"

func main() {
	for i := 1; i < 255; i++ {
		fmt.Printf("Decimal: %d Base Two: %b Hex: %x\n", i, i, i)
	}
}
