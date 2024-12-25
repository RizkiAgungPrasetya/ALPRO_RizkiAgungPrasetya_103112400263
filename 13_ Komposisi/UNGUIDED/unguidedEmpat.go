package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var pita string
	var count int

	// Meminta input jumlah bunga
	fmt.Print("Masukkan jumlah bunga yang ingin dimasukkan: ")
	fmt.Scan(&N)

	for count < N {
		var bunga string
		fmt.Printf("Bunga %d: ", count+1)
		fmt.Scan(&bunga)

		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		if pita == "" {
			pita = bunga // Jika pita kosong, langsung masukkan bunga
		} else {
			pita += " – " + bunga // Gabungkan dengan pemisah
		}
		count++
	}

	// Menampilkan hasil
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", count)
}
