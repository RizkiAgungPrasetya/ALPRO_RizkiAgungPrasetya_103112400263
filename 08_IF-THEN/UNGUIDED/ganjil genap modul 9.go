package main

import (
	"fmt"
)

func main() {
	var bilangan int

	// Meminta input dari pengguna
	fmt.Print("Masukkan suatu bilangan bulat: ")
	fmt.Scan(&bilangan)

	// Mengecek apakah bilangan genap dan negatif
	if bilangan%2 == 0 && bilangan < 0 {
		fmt.Println("Bilangan adalah genap negatif")
	} else {
		fmt.Println("Bilangan bukan genap negatif")
	}
}
