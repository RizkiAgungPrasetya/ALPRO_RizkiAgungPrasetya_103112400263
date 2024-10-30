package main

import "fmt"

func main() {
	var number int

	// Input the number
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Check if number is negative and even
	isNegativeEven := (number < 0) && (number%2 == 0)

	// Print the result
	fmt.Println(isNegativeEven)
}
