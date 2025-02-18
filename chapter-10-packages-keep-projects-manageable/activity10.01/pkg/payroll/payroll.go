package payroll

import (
	"errors"
	"fmt"
)

type Payer interface {
	pay() (string, float64)
}

// payDetails Imprime el nombre y el pago de una int.Pay()
func PayDetails(p Payer) {
	name, pay := p.pay()
	fmt.Printf("%s got paid %.f for the year\n", name, pay)
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
