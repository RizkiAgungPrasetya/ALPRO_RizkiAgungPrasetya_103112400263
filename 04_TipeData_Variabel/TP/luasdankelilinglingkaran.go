package main

import "fmt"

func main() {
	var radius float64
	const pi = 3.14159
	// Meminta input dari pengguna untuk jari-jari lingkaran
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&radius)
	// Menghitung luas lingkaran
	area := pi * radius * radius
	// Menghitung keliling lingkaran
	circumference := 2 * pi * radius
	// Menampilkan hasil
	fmt.Printf("Luas lingkaran: %.2f\n", area)
	fmt.Printf("Keliling lingkaran: %.2f\n", circumference)
}
