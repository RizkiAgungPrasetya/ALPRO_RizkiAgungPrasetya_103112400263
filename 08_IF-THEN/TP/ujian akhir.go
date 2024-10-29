package main

import (
	"fmt"
)

func main() {
	var nilai int

	// Meminta pengguna untuk memasukkan nilai ujian
	fmt.Print("Masukkan nilai ujian: ")
	fmt.Scan(&nilai)

	// Mengecek apakah nilai lulus atau tidak
	if nilai >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
