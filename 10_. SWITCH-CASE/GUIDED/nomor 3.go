package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	var tarif int

	fmt.Scan(&kendaraan)
	fmt.Print("masukan durasi parkir (dalam jam ) :")
	fmt.Scan(&durasi)

	switch {
	case kendaraan == "motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case kendaraan == "motor" && durasi >= 2:
		tarif = 9000
	case kendaraan == "mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case kendaraan == "mobil" && durasi >= 2:
		tarif = 20000
	case kendaraan == "truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case kendaraan == "truk" && durasi >= 2:
		tarif = 35000
	default:
		fmt.Print("jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Print("Tarif Parkir: RP %d\n", tarif)
}
