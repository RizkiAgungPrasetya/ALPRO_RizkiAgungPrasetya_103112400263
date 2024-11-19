package main

import (
	"fmt"
)

func main() {
	// Minta pengguna untuk memasukkan usia
	var usia int
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scan(&usia)

	// Menggunakan switch case untuk menentukan kategori usia
	switch {
	case usia >= 0 && usia <= 12:
		fmt.Println("Kategori Usia: Anak-anak")
	case usia >= 13 && usia <= 17:
		fmt.Println("Kategori Usia: Remaja")
	case usia >= 18 && usia <= 64:
		fmt.Println("Kategori Usia: Dewasa")
	case usia >= 65:
		fmt.Println("Kategori Usia: Lansia")
	default:
		fmt.Println("Usia tidak valid")
	}
}
