package main

import (
	"fmt"
)

func main() {

	var (
		a, b int
	)

	// Inputan untuk memasukan angka pertama dan kedua.
	fmt.Print("Masukan a: ")
	fmt.Scan(&a)

	fmt.Print("Masukan b: ")
	fmt.Scan(&b)

	// Looping yang digunakan untuk menampilkan angka berurut mulai dari inputan a, dan berakhir di inputan b.
	for i := a; i <= b; i++ {
		fmt.Print(i)
	}

}
