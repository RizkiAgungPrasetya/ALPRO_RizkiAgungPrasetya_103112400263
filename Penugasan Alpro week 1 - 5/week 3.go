package main

import (
	"fmt"
)

// Fungsi untuk menentukan apakah mahasiswa cumlaude
func isCumlaude(semester int, ePrT int) bool {
	return semester <= 8 && ePrT >= 500
}

func main() {
	var semester int
	var ePrT int

	// Input jumlah semester dan nilai EPrT
	fmt.Print("Masukkan jumlah semester yang ditempuh: ")
	fmt.Scan(&semester)
	fmt.Print("Masukkan skor EPrT: ")
	fmt.Scan(&ePrT)

	// Memeriksa apakah mahasiswa cumlaude
	if isCumlaude(semester, ePrT) {
		fmt.Println("Mahasiswa lulus dengan predikat cumlaude: true")
	} else {
		fmt.Println("Mahasiswa lulus dengan predikat cumlaude: false")
	}
}
