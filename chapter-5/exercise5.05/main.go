// Package main summing numbers
package main

import "fmt"

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))

}

func sum(nums ...int) (total int) {
	for _, v := range nums {
		total += v
	}
	return
}
