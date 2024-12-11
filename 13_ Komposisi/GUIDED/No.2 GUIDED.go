package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Masukkan tiga bilangan bulat:")
	fmt.Scan(&a, &b, &c)

	// Mencari nilai terbesar
	var max = a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	// Mencari nilai terkecil
	var min = a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	fmt.Printf("Bilangan terbesar: %d\n", max)
	fmt.Printf("Bilangan terkecil: %d\n", min)
}
