package main

import "fmt"

// Fungsi untuk memeriksa apakah angka adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk menghasilkan bilangan prima hingga batas tertentu
func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scan(&n)

	primes := generatePrimes(n)
	fmt.Println("Bilangan prima sampai", n, ":", primes)
}
