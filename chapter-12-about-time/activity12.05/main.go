// Package main Printing the local time in different time zones
package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	nyZoneTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	laZoneTime, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}
	nyTime := current.In(nyZoneTime)
	laTime := current.In(laZoneTime)
	fmt.Println("The local current time is: ", current.Format(time.ANSIC))
	fmt.Println("The time in New York is: ", nyTime.Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is: ", laTime.Format(time.ANSIC))

}
