// Package main Printing the Go representation of a variable
package main

import "fmt"

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	fname := "Joe"
	grades := []int{100, 85, 56}
	states := map[string]string{"KEY": "Kentucky", "VA": "Virginia"}
	p := person{lname: "LIncoln", age: 200, salary: 2333.00}
	fmt.Printf("fname value %#v\n", fname)
	fmt.Printf("grades value %#v\n", grades)
	fmt.Printf("states value %#v\n", states)
	fmt.Printf("p value %#v\n", p)
}
