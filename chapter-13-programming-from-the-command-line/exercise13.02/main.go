// Package main using flags to say hello conditionally
package main

import (
	"flag"
	"fmt"
)

var (
	nameFlag  = flag.String("name", "Sam", "Name of the person to say hello to")
	quiteFlag = flag.Bool("quiet", false, "Toggle to be quiet when saying hello")
)

func main() {
	flag.Parse()
	if !*quiteFlag {
		greeting := fmt.Sprintf("Hello, %s! welcome to the command line.", *nameFlag)
		fmt.Println(greeting)
	} else {
		fmt.Println("hello")
	}
}
