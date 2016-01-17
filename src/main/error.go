package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errorEmptyString = errors.New("Unwilling to print an empty string")
)

func eprinter(msg string) error {
	if msg == "" {
		return fmt.Errorf("eeeee Error!")
	}
	_, err := fmt.Printf("%s\n", msg)

	return err
}

func eprinter2(msg string) error {
	if msg == "" {
		return errorEmptyString
	}

	_, err := fmt.Printf("%s\n", msg)

	return err
}

func main() {
	if err := eprinter("Hello world!"); err != nil {
		fmt.Printf("eprinter faild : %s\n", err)
		os.Exit(1)
	}

	if err := eprinter(""); err != nil {
		fmt.Printf("eprinter faild : %s\n", err)
		//os.Exit(1)
	}

	if err := eprinter2(""); err != nil {
		if err == errorEmptyString {
			fmt.Printf("You tried to print an empty string!\n")
		} else {
			fmt.Printf("eprinter2 faild : %s\n", err)
		}
		os.Exit(1)
	}


}