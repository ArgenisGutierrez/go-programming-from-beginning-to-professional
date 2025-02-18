// Package main calculating pay and performance review
package main

import (
	"fmt"
	"os"

	p "activity10.01/pkg/payroll"
)

var employeeReview = make(map[string]interface{})

func init() {
	println("Welcome to the Employee Pay and Performance Review")
	println("***************************************************")
}

func init() {
	println("Initializing variables")
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
}

func main() {

	d := p.Developer{Individual: p.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}

	m := p.Manager{Individual: p.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.PayDetails(d)
	p.PayDetails(d)
	p.PayDetails(m)
}
