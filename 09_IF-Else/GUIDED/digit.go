package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Masukkan bilangan empat digit: ")
	fmt.Scan(&number)

	// Memeriksa apakah bilangan berada dalam rentang 1000-9999
	if number < 1000 || number > 9999 {
		fmt.Println("Bilangan harus terdiri dari empat digit (1000 - 9999).")
		return
	}

	digit1 := number / 1000
	digit2 := (number / 100) % 10
	digit3 := (number / 10) % 10
	digit4 := number % 10

	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Bilangan membesar.")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		fmt.Println("Bilangan mengecil.")
	} else {
		fmt.Println("Bilangan tidak terurut.")
	}
}