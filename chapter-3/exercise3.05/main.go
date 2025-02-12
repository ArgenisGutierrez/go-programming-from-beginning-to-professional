// Package main safely looping over a string
package main

import "fmt"

func main() {
	logLevel := "デバッグ"
	for i, runeVal := range logLevel {
		fmt.Println(i, string(runeVal))
	}
}
