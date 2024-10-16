package main

import (
	"fmt"
)

func main() {

	var (
		bilSatu, bilDua, hasil int
	)

	// Inputan untuk memasukan berapa kali berapa
	fmt.Print("Masukan bilangan pertama: ")
	fmt.Scan(&bilSatu)

	fmt.Print("Masukan bilangan kedua: ")
	fmt.Scan(&bilDua)

	// Perulangan untuk perkalian tanpa operator kali (*)
	for i := 1; i <= bilDua; i += 1 {
		hasil = hasil + bilSatu
	}

	fmt.Print(hasil)
}
