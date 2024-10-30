package main

import (
	"fmt"
)

func calculateMotorcycles(people int) int {
	// If number of people is even, divide by 2
	// If odd, add 1 motorcycle for the remaining person
	if people%2 == 0 {
		return people / 2
	}
	return (people / 2) + 1
}

func main() {
	var numberOfPeople int

	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&numberOfPeople)

	motorcyclesNeeded := calculateMotorcycles(numberOfPeople)
	fmt.Printf("Jumlah motor yang diperlukan: %d\n", motorcyclesNeeded)
}
