package payroll

import "fmt"

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]any
}

// Pay Imprimer el nombre y el pago del dev
func (d Developer) pay() (string, float64) {
	name := fmt.Sprintf("%s %s", d.Individual.FirstName, d.Individual.LastName)
	pay := d.HourlyRate * d.HoursWorkedInYear
	return name, pay
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
