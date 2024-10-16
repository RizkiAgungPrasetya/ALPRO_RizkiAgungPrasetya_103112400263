package main

import (
	"fmt"
	"math"
)

func main() {

	// Membuat variabel untuk menyimpan nilai inisiasi perulangan dan variabel untuk menyimpan operasi volume.
	var (
		n                        int
		jariJari, tinggi, volume float64
	)

	// Inputan untuk memasukan jumlah perulangan
	fmt.Print("Massukan inisiasi: ")
	fmt.Scan(&n)

	// Perulangan untuk mengulangi operasi mencari volume kerucut
	for i := 1; i <= n; i++ {
		fmt.Println("Masukan jari-jari: ")
		fmt.Scan(&jariJari)

		fmt.Println("Masukan tinggi: ")
		fmt.Scan(&tinggi)

		// Rumus untuk mencari volume kerucut
		volume = (1.0 / 3.0) * math.Pi * jariJari * jariJari * tinggi

		// Menampilkan hasil untuk volume sebuah kerucut.
		fmt.Println("Volume segitiga adalah", volume)
	}

}
