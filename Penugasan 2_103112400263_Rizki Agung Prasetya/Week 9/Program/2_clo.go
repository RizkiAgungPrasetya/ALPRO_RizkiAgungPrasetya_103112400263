package main

import "fmt"

func main() {

	var harga int
	var asessmen bool

	fmt.Scan(&harga, &asessmen)

	if asessmen == true {
		diskon := harga * 35 / 100
		fmt.Print(harga - diskon)
	} else {
		fmt.Println(harga)
	}

}
