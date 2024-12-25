package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&a)

	c = 0

	for b = 0; b <= a; b++ {
		if b%2 != 0 {
			c = c + 1
		}
	}
	fmt.Print("Terdapat ", c, " bilangan ganjil")
}
