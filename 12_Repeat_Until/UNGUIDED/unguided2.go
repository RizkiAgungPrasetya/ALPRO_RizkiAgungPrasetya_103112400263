package main

import (
	"fmt"
)

func main() {
	var input float64
	var roundedUp float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)

	if input == float64(int(input)) {
		roundedUp = input
	} else {
		roundedUp = float64(int(input)) + 1
	}

	current := input

	for current < roundedUp {
		current += 0.1
		fmt.Printf("%.1f\n", current)
	}
}
