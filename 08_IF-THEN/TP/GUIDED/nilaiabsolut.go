package main

import "fmt"

func main() {
	var number int

	// Input the number
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Calculate absolute value
	if number < 0 {
		number = -number
	}

	// Print the result
	fmt.Printf("Absolute value: %d\n", number)
}
