package main

import (
	"fmt"
)

func main() {
	var N int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	// Memastikan input adalah bilangan positif
	if N <= 0 {
		fmt.Println("Silakan masukkan bilangan bulat positif.")
		return
	}

	// Mencetak kuadrat dari 1 hingga N
	fmt.Println("Hasil kuadrat dari bilangan 1 hingga", N, ":")
	for i := 1; i <= N; i++ {
		fmt.Printf("%d ", i*i)
	}
	fmt.Println() // Untuk membuat baris baru setelah output
}
