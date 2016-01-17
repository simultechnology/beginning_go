package main

import (
	"fmt"
	"time"
	"os"
//	"strings"
)

func printer(msg string, msg2 string, repeat int) error {

	var err error

	for repeat > 0 {
		if (err == nil) {
			_, err = fmt.Printf("%s", msg)
		}
		if (err == nil) {
			_, err = fmt.Printf("%s\n", msg2)
		}
		repeat -= 1
	}
	return err
}

func printer2(msg string) (string, error) {
	msg += "\n"

	_, err := fmt.Printf("%s", msg)

	return msg, err
}

func printer3(msg string) error {
	defer fmt.Printf("\n")
	defer fmt.Printf("More\n")

	time.AfterFunc(1 * time.Millisecond, func() {
		fmt.Printf("\nprinter3!\n")
	})

	_, err := fmt.Printf("%s", msg)

	time.Sleep(10 * time.Millisecond)

	return err
}

func printFile(msg string) error {

	f, err := os.Create("helloworld.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write([]byte(msg))
	//bs, _ := strings.NewReader(msg).ReadByte()
	//f.Write(bs)

	return err
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func split2(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

func multiPrinter(format string, msgs ...string)  {

	for _, m := range msgs {
		fmt.Printf(format, m)
	}
}

func main() {

	printer("Hello ", "World!", 5)

	appendedMessage, err := printer2("Hello 石!")
	if err == nil {
		fmt.Printf("%q\n", appendedMessage)
		fmt.Printf("% x\n", appendedMessage)
	}

	printer3("Hello World3!")
	fmt.Printf("end!\n")


	printFile("Hello World33!")

	x, y := split(12)
	fmt.Printf("%d, %d\n", x, y)

	x1, y1 := split2(12)
	fmt.Printf("%d, %d\n", x1, y1)

	multiPrinter("msg : %s\n", "hello, eorld", "this is a pen!", "la pan")
}
