// Package main Formatting a date according to user requirements
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	current := time.Now()
	hour := strconv.Itoa(current.Hour())
	miutes := strconv.Itoa(current.Minute())
	seconds := strconv.Itoa(current.Second())
	year := strconv.Itoa(current.Year())
	month := strconv.Itoa(int(current.Month()))
	day := strconv.Itoa(current.Day())
	fmt.Printf("%s:%s:%s %s/%s/%s\n", hour, miutes, seconds, year, month, day)
}
