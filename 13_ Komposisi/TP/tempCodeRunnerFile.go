package main

import "fmt"

// Fungsi untuk menghitung jumlah faktor dari sebuah bilangan (kecuali bilangan itu sendiri)
func sumOfFactors(n int) int {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

// Fungsi utama untuk memeriksa apakah bilangan merupakan bilangan sempurna
func isPerfectNumber(n int) bool {
	return sumOfFactors(n) == n
}

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&n)

	if isPerfectNumber(n) {
		fmt.Println("Ya, bilangan tersebut adalah bilangan sempurna.")
	} else {
		fmt.Println("Tidak, bilangan tersebut bukan bilangan sempurna.")
	}
}
