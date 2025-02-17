// Package main calculating the wording hours of employees
package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (dev *developer) LogHours(day Weekday, hours int) {
	dev.WorkWeek[day] = hours
}

func (dev developer) HoursWorked() int {
	total := 0
	for _, v := range dev.WorkWeek {
		total += v
	}
	return total
}

func main() {
	d := developer{Individual: Employee{Id: 1, FirstName: "Tony", LastName: "Stark"}, HourlyRate: 10}
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	fmt.Println("Hours worked on Monday:  ", d.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday:  ", d.WorkWeek[Tuesday])
	fmt.Printf("Hours worked this week:  %d\n", d.HoursWorked())
}
