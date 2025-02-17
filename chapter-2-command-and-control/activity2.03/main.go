// Package main Bubble sort
package main

import "fmt"

func bubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println(bubbleSort(nums))
}
