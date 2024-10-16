package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan data
	var a, b, hasil int

	// Inputan untuk memasukan angka dan pangkat
	fmt.Print("Masukan angka: ")
	fmt.Scan(&a)

	fmt.Print("Masukan pangkat: ")
	fmt.Scan(&b)

	// Untuk mendefinisikan nilai dari hasil
	hasil = 1

	// Perulangan untuk membuat operasi perpangkatan
	for i := 1; i <= b; i++ {
		hasil *= a
	}

	// Menampilkan hasil dari operasi perpangkatan
	fmt.Print(hasil)
}
