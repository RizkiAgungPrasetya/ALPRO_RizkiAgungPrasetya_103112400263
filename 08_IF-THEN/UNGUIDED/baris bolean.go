package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat positif y: ")
	fmt.Scan(&y)

	// Menentukan apakah x adalah faktor dari y
	isXFactorOfY := (y%x == 0)

	// Menentukan apakah y adalah faktor dari x
	isYFactorOfX := (x%y == 0)

	// Menampilkan hasil
	fmt.Println(isXFactorOfY)
	fmt.Println(isYFactorOfX)
}
