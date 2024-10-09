package main

import (
	"fmt"
)

func main() {
	var berat, tinggi, hasil float32

	fmt.Println("Masukan berat")
	fmt.Scanln(&berat)
	fmt.Println("Masukan tinggi")
	fmt.Scanln(&tinggi)

	hasil = berat / (tinggi * tinggi)
	fmt.Printf("Jadi hasil BMI nya adalah %.2f", hasil)
}
