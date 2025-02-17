// Package main calculating pay and performance review
package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]any
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

// Pay Imprimer el nombre y el pago del dev
func (d Developer) Pay() (string, float64) {
	name := fmt.Sprintf("%s %s", d.Individual.FirstName, d.Individual.LastName)
	pay := d.HourlyRate * d.HoursWorkedInYear
	return name, pay
}

// Pay Imprime el nombre y el pago del manager
func (m Manager) Pay() (string, float64) {
	name := fmt.Sprintf("%s %s", m.Individual.FirstName, m.Individual.LastName)
	pay := m.Salary + (m.Salary * m.CommissionRate)
	return name, pay
}

// payDetails Imprime el nombre y el pago de una int.Pay()
func payDetails(p Payer) {
	name, pay := p.Pay()
	fmt.Printf("%s got paid %.f for the year\n", name, pay)
}

// ReviewRating Crea un promedio de los review de un dev
func (d Developer) ReviewRating() error {
	total := 0
	for _, inter := range d.Review {
		rating, err := overalReview(inter)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	name := d.Individual.FirstName + " " + d.Individual.LastName
	fmt.Printf("%s got a review reating of %.2f\n", name, averageRating)
	return nil
}

// overalReview Si existen strings reviews los convierte a int
func overalReview(i any) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil

	default:
		return 0, errors.New("uknow type")
	}
}

// convertReviewToInt Convierte los string review a int
func convertReviewToInt(review string) (int, error) {
	switch review {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + review)
	}

}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}

	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d)
	payDetails(m)
}
