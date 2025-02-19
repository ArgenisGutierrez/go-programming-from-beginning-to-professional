// Package main Calculating the future date and time
package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println("the current time: ", current.Format(time.ANSIC))
	duration := time.Duration(21966 * time.Second)
	after := current.Add(duration)
	fmt.Println("6 hours, 6 minutes and 6 seconds from now the time will be: ", after.Format(time.ANSIC))
}
