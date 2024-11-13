package main

import (
	"fmt"
)

func main() {
	var usia int
	var memilikiKTP bool

	// Membaca input usia
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	// Membaca input status kepemilikan KTP
	fmt.Print("Apakah Anda memiliki KTP? (true/false): ")
	fmt.Scan(&memilikiKTP)

	// Menentukan apakah bisa membuat KTP
	if usia >= 17 && memilikiKTP {
		fmt.Println("Anda bisa membuat KTP.")
	} else {
		fmt.Println("Anda tidak bisa membuat KTP.")
	}
}