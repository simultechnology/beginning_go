package main

import (
	"fmt"
	"math/rand"
	"time"
)

func emit(c chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		c <- word
	}

	close(c)
}

func makeRandoms(c chan int)  {
	for i := 0; i < 10; i += 1 {
		c <- rand.Intn(1000)
	}

	close(c)
}

func makeID(idChan chan int) {
	var id int
	id = 0

	for {
		time.Sleep(1000 * time.Millisecond)
		idChan <-id
		id += 1
	}
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)

	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	randoms := make(chan int)

	go makeRandoms(randoms)

	for n := range randoms {
		fmt.Printf("%d ", n)
	}
	fmt.Printf("\n")

	fmt.Printf("----------------------------------------\n")
	idChan := make(chan int)

	go makeID(idChan)

	fmt.Printf("id : %d\n", <-idChan)
	fmt.Printf("id : %d\n", <-idChan)
	fmt.Printf("id : %d\n", <-idChan)
	fmt.Printf("id : %d\n", <-idChan)
}