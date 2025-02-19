// Package main duration of execution
package main

import (
	"fmt"
	"strconv"
	"time"
)

func elapsedTime(start, end time.Time) string {
	elapsed := end.Sub(start)
	hours := strconv.Itoa(int(elapsed.Hours()))
	minutes := strconv.Itoa(int(elapsed.Minutes()))
	seconds := strconv.Itoa(int(elapsed.Seconds()))
	return fmt.Sprintf("The total execution time elapsed is: %s hour(s) and %s minute(s) and %s seconds", hours, minutes, seconds)
}

func main() {
	start := time.Now()
	time.Sleep(5 * time.Second)
	end := time.Now()
	fmt.Println(elapsedTime(start, end))

}
