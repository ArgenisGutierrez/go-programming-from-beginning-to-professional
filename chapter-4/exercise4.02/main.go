package main

import "fmt"

func compArrays() (bool, bool, bool) {
	var arra1 [5]int
	arra2 := [5]int{0}
	arra3 := [...]int{0, 0, 0, 0, 0}
	arra4 := [5]int{0, 0, 0, 0, 9}
	return arra1 == arra2, arra1 == arra3, arra1 == arra4
}

func main() {
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0} :", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)
	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3)
}
