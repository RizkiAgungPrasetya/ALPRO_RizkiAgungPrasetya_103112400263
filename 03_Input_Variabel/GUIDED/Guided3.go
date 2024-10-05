package main

import "fmt"

// Fungsi untuk menghitung luas segitiga
func hitungLuasSegitiga(alas float64, tinggi float64) float64 {
	return 0.5 * alas * tinggi
}

func main() {
	var alas float64
	var tinggi float64

	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scanln(&alas)

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)

	luas := hitungLuasSegitiga(alas, tinggi)
	fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, luas)
}
