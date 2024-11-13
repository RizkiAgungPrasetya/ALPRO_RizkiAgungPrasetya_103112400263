package main

import "fmt"

func main() {
	var (
		jari  float32
		rumus float32
	)
	fmt.Print("Masukan jari-jari: ")
	fmt.Scanln(&jari)
	rumus = jari * jari * 3.14
	fmt.Println("Luas lingkaran adalah", rumus)
}
