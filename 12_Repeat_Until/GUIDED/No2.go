package main

import "fmt"

func main() {
	var bilangan int

	// Pengulangan akan terus berjalan sampai mendapat bilangan positif
	for {
		fmt.Print("Masukkan sebuah bilangan: ")
		fmt.Scanln(&bilangan)

		// Cek apakah bilangan positif
		if bilangan > 0 {
			fmt.Printf("%d adalah bilangan bulat positif\n", bilangan)
			break
		} else {
			fmt.Println("Bukan bilangan bulat positif, silakan coba lagi!")
		}
	}
}
