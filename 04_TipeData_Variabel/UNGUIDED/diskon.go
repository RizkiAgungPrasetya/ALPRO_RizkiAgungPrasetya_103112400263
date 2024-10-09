package main

import (
	"fmt"
)

func main() {
	var harga, diskon, total int32

	fmt.Println("Masukan total harga:")
	fmt.Scanln(&harga)
	fmt.Println("Masukan diskon:")
	fmt.Scanln(&diskon)

	total = harga - (diskon * harga / 100)
	fmt.Println("Total belanjaan adalah ", total)
}
