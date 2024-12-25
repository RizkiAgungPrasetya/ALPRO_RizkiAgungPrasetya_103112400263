package main

import (
	"fmt"
)

func main() {
	var target, donation, totalDonations, donorCount int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	for totalDonations < target {
		donorCount++
		fmt.Printf("Masukkan donasi dari donatur %d: ", donorCount)
		fmt.Scan(&donation)

		totalDonations += donation

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donorCount, donation, totalDonations)
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonations, donorCount)
}
