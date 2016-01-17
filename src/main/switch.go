package main

import (
	"fmt"
	"os"
)

func main() {

	n, err := fmt.Printf("Hello, World!!")

	n = 0

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output")
		fallthrough
	case n != 14:
		fmt.Printf("Wrong number of bytes: %d", n)
	default:
		fmt.Printf("OK!")
	}
	fmt.Printf("\n")

	atoz := "The quick brown fox jumps over the lazy dog"

	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'i', 'u', 'e', 'o':
			vowels += 1
		case 'z':
			zeds += 1
			fallthrough
		default:
			consonants += 1
		}
	}

	fmt.Printf("vowels : %d, consonants : %d, zeds : %d\n", vowels, consonants, zeds)

	fmt.Println(string("石"[0:]))
}
