// Package main enforcing a specific format of date and time
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	date := time.Date(2025, 2, 19, 12, 18, 20, 29399381, time.UTC)
	day := strconv.Itoa(date.Day())
	month := strconv.Itoa(int(date.Month()))
	year := strconv.Itoa(date.Year())
	hour := strconv.Itoa(date.Hour())
	minute := strconv.Itoa(date.Minute())
	second := strconv.Itoa(date.Second())
	fmt.Println(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day)

}
