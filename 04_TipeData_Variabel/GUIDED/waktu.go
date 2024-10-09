package main

import (
	"fmt"
)

func main() {
	var m, jam, menit, detik int32

	fmt.Println("Masukan total waktu")
	fmt.Scanln(&m)

	jam = m / 3600
	menit = (m % 3600) / 60
	detik = (m % 3600) % 60

	fmt.Print(jam, " Jam ", menit, " Menit ", detik, " Detik ")
}
