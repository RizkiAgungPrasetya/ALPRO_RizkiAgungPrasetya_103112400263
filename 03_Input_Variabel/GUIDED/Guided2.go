package main

import (
	"fmt"
	"math"
)

func main() {
	var fx float64

	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	// Mengisolasi x: x^2/2 = fx - 5
	xKuadrat := 2 * (fx - 5)

	// Menghitung akar kuadrat (nilai mutlak x)
	nilaiMutlakX := math.Sqrt(xKuadrat)

	// Menampilkan hasil
	fmt.Printf("Nilai x adalah: %.2f atau %.2f\n", nilaiMutlakX, -nilaiMutlakX)
}
