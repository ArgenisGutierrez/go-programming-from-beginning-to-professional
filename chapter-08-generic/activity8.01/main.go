// Package main a minimum value
package main

import "fmt"

func findMinGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	for _, num := range nums {
		if num < max {
			max = num
		}
	}
	return max
}

func main() {
	minInt := findMinGeneric([]int{1, 32, 5, 8, 10, 11})
	fmt.Println("Min value: ", minInt)
	minFloat := findMinGeneric([]float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1})
	fmt.Println("Min value: ", minFloat)
}
