package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan data
	var a, hasil int

	// Inputan untuk memasukan bilangan bulat
	fmt.Print("Masukan angka: ")
	fmt.Scan(&a)

	// Untuk mendefinisikan nilai dari hasil
	hasil = 1

	// Perulangan untuk membuat operasi faktorial
	for i := 1; i <= a; i++ {
		hasil *= i
	}

	// Menampilkan hasil dari faktorial
	fmt.Print(hasil)
}
