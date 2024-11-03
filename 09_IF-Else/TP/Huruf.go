package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// Meminta input dari pengguna
	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scan(&input)

	// Mengubah input menjadi huruf kapital untuk mempermudah pemeriksaan
	input = strings.ToUpper(input)

	// Memeriksa apakah input adalah huruf vokal
	if len(input) == 1 && (input == "A" || input == "I" || input == "U" || input == "E" || input == "O") {
		fmt.Println("Huruf Vokal")
	} else if len(input) == 1 && (input >= "B" && input <= "Z") { // Memeriksa apakah input adalah huruf konsonan
		fmt.Println("Huruf Konsonan")
	} else {
		fmt.Println("Input tidak valid. Harap masukkan satu huruf.")
	}
}
