// Package main Omitiendo partes en la declaracion de variables
package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool
	LogLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
