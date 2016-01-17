package main

import (
	"fmt"
	"myutils"
)

const (
	message = "The answer to life is %d\n"
	answer  = 42
)

const (
	message1 = "%d %d\n"
	answer1 = iota
	answer2
)

func main() {
	fmt.Println("start!")
	fmt.Printf("%v\n", myutils.Sum(3, 5))
	fmt.Printf("Hello, %s\n", myutils.Message)

	fmt.Printf(message, answer)
	fmt.Printf(message1, answer1, answer2)

	pi := 3.14
	fmt.Printf("value : %.2f\n", pi)

	nine := 2000
	fmt.Printf("%v\n", nine)

	i := uint8(255)
	fmt.Printf("%v\n", i)

	b := byte(64)
	fmt.Printf("%x\n", b)
}
