package main

import (
	"fmt"
	"time"
	"os"
)

func main() {

	//counter := 0

	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("Timeout")
		os.Exit(0)
	})

	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {

//		t := time.Now().Hour()
//
//		fmt.Printf("Hello World! %v\n", t)
		//counter += 1
		//fmt.Println(counter)
		fmt.Printf("Hello World! %v\n", j)
	}

}
