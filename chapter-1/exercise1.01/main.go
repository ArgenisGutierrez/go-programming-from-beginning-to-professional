package main

import "fmt"

type Patient struct {
	name    string
	lName   string
	age     int
	allergy bool
}

func main() {
	patient := Patient{
		name:    "Bob",
		lName:   "Smith",
		age:     34,
		allergy: false,
	}

	fmt.Println(patient.name)
	fmt.Println(patient.lName)
	fmt.Println(patient.age)
	fmt.Println(patient.allergy)

}
