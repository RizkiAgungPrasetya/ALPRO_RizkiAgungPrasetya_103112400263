package main

import "fmt"

// Fungsi untuk menghitung volume kubus
func hitungKubus(sisi float64) float64 {
	return sisi * sisi * sisi
}

func main() {
	var sisi float64
	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scanln(&sisi)

	volume := hitungKubus(sisi)
	fmt.Printf("Volume kubus dengan sisi %.2f adalah %.2f\n", sisi, volume)
}
