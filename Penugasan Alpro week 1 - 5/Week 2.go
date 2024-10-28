package main

import (
	"fmt"
)

func main() {
	var p, l int

	// Input panjang dan lebar
	fmt.Print("Masukkan panjang (p): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar (l): ")
	fmt.Scan(&l)

	// Menghitung luas dan keliling
	Z := p * l       // Luas
	K := 2 * (p + l) // Keliling

	// Menampilkan hasil
	fmt.Printf("Keliling (K): %d\n", K)
	fmt.Printf("Luas (Z): %d\n", Z)
}
