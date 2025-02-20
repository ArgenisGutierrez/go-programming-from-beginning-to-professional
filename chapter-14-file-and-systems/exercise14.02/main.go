// Package main backing up files
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	ErrWorkingFileNotFound = errors.New("the working file is not found")
)

func createBakup(working, backup string) error {
	// Check to see if our file exists
	_, err := os.Stat(working)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrWorkingFileNotFound
		}
		return err
	}
	// Open the file
	workFile, err := os.Open(working)
	if err != nil {
		return err
	}
	// Read the file content
	content, err := io.ReadAll(workFile)
	if err != nil {
		return err
	}
	// Create the back up
	err = os.WriteFile(backup, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func addNotes(workingFile, notes string) error {
	// Add jump line to note
	notes += "\n"
	// Open the file
	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	// Add the note
	if _, err := f.Write([]byte(notes)); err != nil {
		return err
	}
	return nil
}

func main() {
	backupFile := "backupFile.txt"
	workingFile := "note.txt"
	data := "note"
	err := createBakup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 10; i++ {
		note := data + " " + strconv.Itoa(i)
		err := addNotes(workingFile, note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
