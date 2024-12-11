package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor-faktor dari %d adalah: ", n)

	// Mencari faktor-faktor
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i)
			// Menambahkan spasi jika bukan faktor terakhir
			if i != n {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
}
