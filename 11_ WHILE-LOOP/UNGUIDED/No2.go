package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&input)

	for input > 0 {
		nomer := input % 10
		fmt.Print(nomer)
		input /= 10
	}
}
