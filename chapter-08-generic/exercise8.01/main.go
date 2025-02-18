// Package main calculate the maximum value using interfaces
package main

import "fmt"

type Number interface {
	int | float64
}

func findMaxGeneric[Num Number](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	maxGenericInt := findMaxGeneric([]int{1, 32, 6, 8, 19, 11})
	fmt.Printf("Max generic int: %v\n", maxGenericInt)
	maxFloat := findMaxGeneric([]float64{1.1, 32.1, 5.1, 6.1, 7.1, 19.1, 12.1})
	fmt.Printf("Max generic float : %v\n", maxFloat)

}
