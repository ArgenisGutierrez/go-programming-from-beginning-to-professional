// Package main creating a function to return a timestamp
package main

import (
	"fmt"
	"time"
)

func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}

func main() {
	fmt.Println(whatstheclock())
}
