package main

import "fmt"

func main() {

	// Membuat variable
	var ph float32

	// Meminta untuk memasukan ph air
	fmt.Print("Masukkan pH: ")
	fmt.Scan(&ph)

	// Switch case untuk menentukan apakah air layak minum
	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air Layak Minum")
	case (ph < 6.5 || ph > 8.6) && ph <= 14:
		fmt.Println("Air Tidak Layak Minum")
	default:
		fmt.Print("Input tidak valid, rentang pH 0 - 14")
	}
}
