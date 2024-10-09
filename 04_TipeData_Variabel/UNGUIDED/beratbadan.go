package main

import (
	"fmt"
)

func main() {
	var berat, tinggi, bmi float32

	fmt.Println("Masukan nilai BMI")
	fmt.Scanln(&bmi)
	fmt.Println("Masukan tinggi")
	fmt.Scanln(&tinggi)
	berat = bmi * tinggi * tinggi
	fmt.Printf("Jadi berat badan akhir adalah %.0f", berat)
}
