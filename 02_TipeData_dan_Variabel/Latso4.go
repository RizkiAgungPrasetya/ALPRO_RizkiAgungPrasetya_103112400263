package main

import "fmt"

func main() {
	var (
		suhu  float32
		rumus float32
	)
	fmt.Print("Masukan suhu farenhit: ")
	fmt.Scanln(&suhu)
	rumus = (suhu - 32) * 5 / 9
	fmt.Println("hasil fareheit ke celcius adalah", rumus)
}
