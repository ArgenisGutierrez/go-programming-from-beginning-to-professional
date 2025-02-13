// Package main Crating a locale checker
package main

import (
	"fmt"
	"os"
	"strings"
)

type Locale struct {
	lenague string
	region  string
}

func getLocales() map[Locale]struct{} {
	supportedLocales := make(map[Locale]struct{}, 5)
	supportedLocales[Locale{"en", "US"}] = struct{}{}
	supportedLocales[Locale{"en", "CN"}] = struct{}{}
	supportedLocales[Locale{"fr", "CN"}] = struct{}{}
	supportedLocales[Locale{"fr", "FR"}] = struct{}{}
	supportedLocales[Locale{"ur", "RU"}] = struct{}{}
	return supportedLocales
}

func localeExist(l Locale) bool {
	_, exist := getLocales()[l]
	return exist
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No locale passed")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := Locale{
		lenague: localeParts[0],
		region:  localeParts[1],
	}

	if !localeExist(passedLocale) {
		fmt.Printf("Locales not Supported: %v\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Locales passed is Supported")
}
