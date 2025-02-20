// Package main Building a program to validate Social Security Numbers
package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	ErrInvalidSSNLength  = errors.New("invalid SSN length")
	ErrInvalidSSNumbers  = errors.New("have non-numeric digits")
	ErrInvalidSSNPrefix  = errors.New("has three zeros as prefix")
	ErrInvalidDigidPlace = errors.New("start with 9 or 9 in the four place")
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}
	log.Printf("Checking data %#v", validateSSN)
	for idx, ssn := range validateSSN {
		log.Printf("Validate Data %#v %d of %d ", ssn, idx+1, len(validateSSN))
		ssn = strings.Replace(ssn, "-", "", -1)
		err := isNumber(ssn)
		if err != nil {
			log.Print(err)
		}
		err = validateLength(ssn)
		if err != nil {
			log.Print(err)
		}
		err = isPrefixValid(ssn)
		if err != nil {
			log.Print(err)
		}
		err = validDigitPlace(ssn)
		if err != nil {
			log.Print(err)
		}
	}

}

func validateLength(ssn string) error {
	ssn = strings.TrimSpace(ssn)
	if len(ssn) != 9 {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNLength)
	}
	return nil
}

func isNumber(ssn string) error {
	_, err := strconv.Atoi(ssn)
	if err != nil {
		return fmt.Errorf("the value of %s caused an erros: %v\n", ssn, ErrInvalidSSNumbers)
	}
	return nil
}

func isPrefixValid(ssn string) error {
	if strings.HasPrefix(ssn, "000") {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidSSNPrefix)
	}
	return nil
}

func validDigitPlace(ssn string) error {
	if string(ssn[0]) == "9" && (string(ssn[3]) != "9" && string(ssn[3]) != "7") {
		return fmt.Errorf("the value of %s caused an error: %v\n", ssn, ErrInvalidDigidPlace)
	}
	return nil
}
