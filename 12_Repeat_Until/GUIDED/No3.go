package main

import "fmt"

func main() {
	var x, y int

	// Input bilangan x dan y
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scanln(&y)

	// Menyimpan nilai awal x untuk output akhir
	bilanganAwal := x

	// Melakukan pengurangan berulang
	for x >= y {
		fmt.Printf("%d\n", x)
		x = x - y
	}
	fmt.Printf("%d\n", x)

	// Pengecekan apakah x kelipatan y
	isKelipatan := x == 0
	fmt.Printf("%d kelipatan %d: %t\n", bilanganAwal, y, isKelipatan)
}
