package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan banyaknya persegi: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var sisi float64
		fmt.Printf("Masukkan sisi persegi ke-%d: ", i+1)
		fmt.Scan(&sisi)

		// Menghitung luas dan keliling
		luas := sisi * sisi
		keliling := 4 * sisi

		// Menampilkan hasil
		fmt.Printf("Persegi ke-%d: Luas = %.2f, Keliling = %.2f\n", i+1, luas, keliling)
	}
}
