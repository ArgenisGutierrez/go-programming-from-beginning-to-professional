// Package main creating various functions to calculate salary
package main

import "fmt"

func main() {
  devSalary := salary(50,2080,developerSalary)
  bossSalary := salary(150000,2500,managerSalary)
  fmt.Printf("Boss salary: %d\n",bossSalary)
  fmt.Printf("Dev salary: %d\n",devSalary)
}

func salary(x,y int,f func(int,int)int) int  {
  pay := f(x,y)
  return pay
}

func managerSalary(baseSalary, bouns int) int {
  return baseSalary + bouns
}

func developerSalary(hourlyRate, hoursWorked int) int {
  return hourlyRate * hoursWorked
}
