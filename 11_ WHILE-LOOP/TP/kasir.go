package main

import (
	"fmt"
)

func main() {
	var total float64
	var item string // Deklarasi variabel item

	for { // While-loop simulasi
		fmt.Print("Masukkan nama barang (atau ketik 'selesai' untuk keluar): ")
		fmt.Scanln(&item)

		if item == "selesai" {
			break
		}

		var price float64
		fmt.Print("Masukkan harga barang: ")
		_, err := fmt.Scanln(&price) // Menangkap error saat input harga
		if err != nil {
			fmt.Println("Input harga tidak valid. Silakan coba lagi.")
			// Mengabaikan input yang tidak valid dan melanjutkan ke iterasi berikutnya
			continue
		}

		total += price
	}

	fmt.Printf("Total belanja: %.2f\n", total)
	fmt.Println("Transaksi selesai. Terima kasih!")
}
