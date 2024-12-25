package main

import (
	"fmt"
)

func main() {
	var number int
	var digitCount int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	for number > 0 {
		number /= 10
		digitCount++
	}

	fmt.Printf("Banyaknya digit: %d\n", digitCount)
}
