package main

import (
	"fmt"
)

func main()  {
	atoz := "石川です The quick brown fox jumps over the lazy dog\n"

	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("%s\n", atoz[16:19])

	for i, r := range atoz {

		//fmt.Printf("%v %v\n", i, string(r))
		if string(r) == " " {
			fmt.Printf("\n%d : ", i)
		} else {
			if i == 0 {
				fmt.Printf("%d : ", i)
			}
			//fmt.Printf("%s", string(r))
			fmt.Printf("%c", r)
		}

	}

	for i := 0; i < len(atoz); i++ {

		//fmt.Printf("%v %v\n", i, string(r))
		if string(atoz[i]) == " " {
			fmt.Printf("\n%d : ", i)
		} else {
			if i == 0 {
				fmt.Printf("%d : ", i)
			}
			fmt.Printf("%s", string(atoz[i]))
		}

	}

	const nihongo = "ほげふが㟋㟌気持ち⻤鬼岩城⻲亀仙人"

	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	fmt.Printf("%d\n", len(atoz))
	fmt.Printf("%d\n", len(nihongo))

	fmt.Printf("%s\n", atoz)

	atoz2 := `石川です The quick brown fox jumps over the lazy dog\n`
	fmt.Printf("%s\n", atoz2)

	for i, r := range atoz {

		//fmt.Printf("%v %v\n", i, string(r))
		if r == ' ' {
			fmt.Printf("\n%d : ", i)
		} else {
			if i == 0 {
				fmt.Printf("%d : ", i)
			}
			//fmt.Printf("%s", string(r))
			fmt.Printf("%c", r)
		}

	}
}
