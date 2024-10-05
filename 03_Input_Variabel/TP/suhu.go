package main

import "fmt"

func main() {
	fmt.Println("Konverter Suhu Fahrenheit ke Kelvin")
	fmt.Println("-----------------------------------")

	var fahrenheit float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	kelvin := (fahrenheit-32)*5/9 + 273.15

	fmt.Printf("%.2f°F sama dengan %.2fK\n", fahrenheit, kelvin)
}
