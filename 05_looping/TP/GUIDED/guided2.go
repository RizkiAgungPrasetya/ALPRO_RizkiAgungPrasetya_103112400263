package main

import (
	"fmt"
)

func main() {

	var (
		inisiasi, alas, tinggi, luas int
	)

	// Inputan untuk mendeklarasikan berapa kalii perulangan akan berlanjut.
	fmt.Print("Massukan inisiasi: ")
	fmt.Scan(&inisiasi)

	// Perulangan untuk mengulangi operasi mencari luas
	for i := 1; i <= inisiasi; i++ {
		fmt.Println("Masukan alas: ")
		fmt.Scan(&alas)

		fmt.Println("Masukan tinggi: ")
		fmt.Scan(&tinggi)

		luas = alas * tinggi / 2
		fmt.Println("Luas segitiga adalah", luas)
	}

}
