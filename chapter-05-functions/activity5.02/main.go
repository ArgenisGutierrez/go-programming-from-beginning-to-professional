// Package main calculating the payable amount for employees based on working hours
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

// HoursWorked Calcula el total de horas trabajadas
func (dev developer) HoursWorked() int {
	total := 0
	for _, v := range dev.WorkWeek {
		total += v
	}
	return total
}

func nonLoggedHours() func(int)int{
  hours:=0
  return func(i int) int {
     hours +=i
     return hours
  }
}

func (dev developer) PayDay() (int,bool) {
  if dev.HoursWorked() > 40{
    extraHours := dev.HoursWorked() - 40
    overTime := extraHours * 2 * dev.HourlyRate
    regularPay := 40 * dev.HourlyRate
    return regularPay + overTime,true
  }
    return dev.HoursWorked() * dev.HourlyRate,false
}

func (dev  developer) PayDetails()  {
  for i, v := range dev.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}
	fmt.Printf("\nHours worked this week:  %d\n", dev.HoursWorked())
	pay, overtime := dev.PayDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
	fmt.Println()
}

func main() {
  d := developer{Individual: Employee{Id: 1, FirstName: "Tony", LastName: "Stark"}, HourlyRate: 10}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	d.LogHours(Wednesday, 10)
	d.LogHours(Thursday, 10)
	d.LogHours(Friday, 6)
	d.LogHours(Saturday, 8)
	d.PayDetails()
	
}
