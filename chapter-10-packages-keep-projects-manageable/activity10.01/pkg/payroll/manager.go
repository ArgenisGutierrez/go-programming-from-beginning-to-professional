package payroll

import "fmt"

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

// Pay Imprime el nombre y el pago del manager
func (m Manager) pay() (string, float64) {
	name := fmt.Sprintf("%s %s", m.Individual.FirstName, m.Individual.LastName)
	pay := m.Salary + (m.Salary * m.CommissionRate)
	return name, pay
}
