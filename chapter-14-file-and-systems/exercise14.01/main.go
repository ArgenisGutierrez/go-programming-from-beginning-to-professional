// Package main simulating a cleanup
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Make a channel used to receive the notifications from the Notify method
	sigs := make(chan os.Signal, 1)
	// Make a channel used to let us know when the program can exit
	done := make(chan struct{})
	// This method send values of the os.Signal type to a channel
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGSTOP)
	// Create a go routine with an infinite loop searching for signal
	go func() {
		for {
			// var receive the signal
			s := <-sigs
			// Action depends of the signal
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted. Someone might have pressed CTRL-C")
				fmt.Println("Some clean up is occuring")
				cleanup()
				// Indicate the exit of the program
				done <- struct{}{}
			case syscall.SIGSTOP:
				fmt.Println()
				fmt.Println("Someone pressed CTRL-Z")
				fmt.Println("Some clean up is occuring")
				cleanup()
				// Indicate the exit of the program
				done <- struct{}{}
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught(ctrl-z, ctrl-c)")
	done <- struct{}{}
	fmt.Println("out of here")
}

func cleanup() {
	fmt.Println("simulating clean up")
	for i := 0; i <= 10; i++ {
		fmt.Println("Deleting Files.. Not really.", i)
		time.Sleep(1 * time.Second)
	}
}
