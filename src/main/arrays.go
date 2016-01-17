package main

import (
	"fmt"
)

func arrayPrinter(w []string) {

	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
	//w[5] = "hoge"
}

func main() {

	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	//fmt.Printf("%T", words)

	arrayPrinter(words)
	fmt.Printf("%v\n", words)
	fmt.Printf("cap : %d\n", cap(words))

	sliceWords := make([]string, 4)
	fmt.Printf("len : %d, cap : %d\n", len(sliceWords), cap(sliceWords))
	sliceWords[0] = "aaaa"
	sliceWords[1] = "nnnn"
	sliceWords[2] = "fuga"
	sliceWords[3] = "hoge"
	fmt.Printf("len : %d, cap : %d\n", len(sliceWords), cap(sliceWords))

	arrayPrinter(sliceWords)

	sliceWords2 := make([]string, 0, 4)
	fmt.Printf("len : %d, cap : %d\n", len(sliceWords2), cap(sliceWords2))
	sliceWords2 = append(sliceWords2, "aishi")
	sliceWords2 = append(sliceWords2, "aishi")
	sliceWords2 = append(sliceWords2, "aishi")
	sliceWords2 = append(sliceWords2, "aishi")
	fmt.Printf("len : %d, cap : %d\n", len(sliceWords2), cap(sliceWords2))
	sliceWords2 = append(sliceWords2, "aishi")
	fmt.Printf("len : %d, cap : %d\n", len(sliceWords2), cap(sliceWords2))
	sliceWords2 = append(sliceWords2, "aishi")
	sliceWords2 = append(sliceWords2, "aishi")
	sliceWords2 = append(sliceWords2, "aishi")
	fmt.Printf("len : %d, cap : %d\n", len(sliceWords2), cap(sliceWords2))
	fmt.Printf("value : %s\n", sliceWords2[7])
	arrayPrinter(sliceWords2)

	newWords := make([]string, 8)
	copy(newWords, sliceWords2)
	newWords[5] = "takaka"
	arrayPrinter(newWords)
	arrayPrinter(sliceWords2)
}
