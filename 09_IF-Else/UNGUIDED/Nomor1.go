package main

import (
	"fmt"
)

func hitungBiayaKirim(berat int) int {
	// Menghitung total berat dalam kg dan sisa berat dalam gram
	beratKg := berat / 1000
	sisaBerat := berat % 1000

	// Biaya dasar per kg
	biayaPerKg := 10000
	totalBiaya := beratKg * biayaPerKg

	// Menentukan biaya tambahan berdasarkan sisa berat
	if beratKg > 10 {
		// Jika total berat lebih dari 10kg, sisa berat digratiskan
		return totalBiaya
	}

	if sisaBerat >= 500 {
		totalBiaya += sisaBerat * 5
	} else {
		totalBiaya += sisaBerat * 15
	}

	return totalBiaya
}

func main() {
	var berat int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&berat)

	if berat < 0 {
		fmt.Println("Berat tidak boleh negatif.")
		return
	}

	biaya := hitungBiayaKirim(berat)
	fmt.Printf("Biaya pengiriman untuk berat %d gram adalah Rp. %d\n", berat, biaya)
}
