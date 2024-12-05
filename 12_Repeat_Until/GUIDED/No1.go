package main

import "fmt"

func main() {
	var kata string
	var jumlahUlang int

	// Meminta input kata dari pengguna
	fmt.Print("Masukkan kata yang ingin diulang: ")
	fmt.Scanln(&kata)

	// Meminta jumlah pengulangan
	fmt.Print("Masukkan jumlah pengulangan yang diinginkan: ")
	fmt.Scanln(&jumlahUlang)

	// Inisialisasi counter
	counter := 1

	// Melakukan pengulangan sampai counter melebihi jumlahUlang
	for {
		fmt.Printf("%d. %s\n", counter, kata)
		counter++

		// Kondisi untuk keluar dari pengulangan
		if counter > jumlahUlang {
			break
		}
	}
}
