package main

import (
	"fmt"
)

// Struktur untuk menyimpan data BMI
type BMIRecord struct {
	Name     string
	Weight   float64 // dalam kg
	Height   float64 // dalam meter
	BMIValue float64
}

// Fungsi untuk menghitung BMI
func calculateBMI(weight float64, height float64) float64 {
	if height == 0 {
		return 0 // Menghindari pembagian dengan nol
	}
	return weight / (height * height)
}

func main() {
	var name string
	var weight, height float64

	// Input data dari pengguna
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scanln(&weight)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scanln(&height)

	// Menghitung BMI
	bmiValue := calculateBMI(weight, height)

	// Membuat record BMI
	bmiRecord := BMIRecord{
		Name:     name,
		Weight:   weight,
		Height:   height,
		BMIValue: bmiValue,
	}

	// Menampilkan hasil
	fmt.Printf("Data BMI:\n")
	fmt.Printf("Nama: %s\n", bmiRecord.Name)
	fmt.Printf("Berat Badan: %.2f kg\n", bmiRecord.Weight)
	fmt.Printf("Tinggi Badan: %.2f m\n", bmiRecord.Height)
	fmt.Printf("Nilai BMI: %.2f\n", bmiRecord.BMIValue)
}
