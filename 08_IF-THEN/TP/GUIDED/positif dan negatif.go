package main

import "fmt"

func main() {
	var number int

	// Input the number
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Check if number is positive
	if number > 0 {
		fmt.Println("positif")
	} else {
		fmt.Println("bukan positif")
	}
}
