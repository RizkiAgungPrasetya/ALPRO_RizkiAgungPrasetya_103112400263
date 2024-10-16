package main

import (
	"fmt"
)

func main() {

	var (
		n, hasil int
	)

	// Inputan untuk memasukan nilai dari n.
	fmt.Print("Masukan nilai n: ")
	fmt.Scan(&n)

	// Perulangan untuk menjumlahkan jumlah deret dari nilai n yang telah dimasukann.
	for i := 1; i <= n; i++ {
		hasil = hasil + i
	}

	fmt.Print(hasil)
}
