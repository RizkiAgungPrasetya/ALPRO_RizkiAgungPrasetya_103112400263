package main

import "fmt"

func main() {

	var bilangan int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&bilangan)

	if bilangan >= 0 {
		fmt.Print("Bilangan Positif")

		if bilangan%2 == 0 {
			fmt.Print(" Genap")
		} else {
			fmt.Print(" Ganjil")
		}
	} else {
		fmt.Println("Bilangan Negatif")
	}
}
