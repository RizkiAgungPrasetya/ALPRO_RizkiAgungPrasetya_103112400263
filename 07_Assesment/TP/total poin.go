package main

import (
	"fmt"
)

func hitungPoin(jumlahBarang int) int {
	// Setiap barang memberi 10 poin
	poinDasar := 10
	totalPoin := jumlahBarang * poinDasar

	// Jika lebih dari 5 barang, tambahkan 5 poin untuk setiap barang setelah barang kelima
	if jumlahBarang > 5 {
		totalPoin += (jumlahBarang - 5) * 5
	}

	return totalPoin
}

func main() {
	var jumlahBarang int

	// Meminta input dari pengguna
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	// Menghitung dan menampilkan total poin
	totalPoin := hitungPoin(jumlahBarang)
	fmt.Printf("Total poin yang didapat: %d\n", totalPoin)
}
