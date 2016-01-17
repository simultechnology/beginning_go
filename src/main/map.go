package main

import (
	"fmt"
)

func main() {
	dayMonth := make(map[string]int)
	dayMonth["Jan"] = 31
	dayMonth["Feb"] = 28
	dayMonth["Mar"] = 31
	dayMonth["Apr"] = 30
	dayMonth["May"] = 31
	dayMonth["Jun"] = 30
	dayMonth["Jul"] = 31
	dayMonth["Aug"] = 31
	dayMonth["Sep"] = 30
	dayMonth["Oct"] = 31
	dayMonth["Nov"] = 30
	dayMonth["Dec"] = 31

	days, ok := dayMonth["Jan"]
	if !ok {
		fmt.Printf("Can't get days for Janualy")
	} else {
		fmt.Printf("%d\n", days)
	}

	daysYear := 0
	for month, day := range dayMonth {
		fmt.Printf("%s has %d days!\n", month, day)
		daysYear += day
	}
	fmt.Printf("year has %d days!\n", daysYear)

	dayMonth2 := map[string]int{
		"Jan!" : 31,
		"Feb!" : 28,
		"Mar!" : 31,
		"Apr!" : 30,
		"May!" : 31,
		"Jun!" : 30,
		"Jul!" : 31,
		"Aug!" : 31,
		"Sep!" : 30,
		"Oct!" : 31,
		"Nov!" : 30,
		"Dec!" : 31,
	}

	has28 := 0

	for month, day := range dayMonth2 {
		if day == 28 {
			has28 += 1
			fmt.Printf("%s has %d!\n", month, day)
		}
	}
	fmt.Printf("how many month have 28 days in one month? It's %d month!\n", has28)

	addOneElem(dayMonth2)

	fmt.Printf("%v\n", dayMonth2)

}

func addOneElem(dayMonth map[string]int) {
	dayMonth["hoge"] = 31
}