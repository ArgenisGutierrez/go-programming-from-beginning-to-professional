// Package main triggering number wraparound
package main

import "fmt"

func main() {
	var (
		a int8  = 125
		b uint8 = 253
	)
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8 ", a, "unit8 ", b)
	}
}
