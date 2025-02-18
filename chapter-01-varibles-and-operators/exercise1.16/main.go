// Package main pointer value swap
package main

import "fmt"

func swap(a, b *int) {
	res := *a
	*a = *b
	*b = res
}
func main() {
	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}
