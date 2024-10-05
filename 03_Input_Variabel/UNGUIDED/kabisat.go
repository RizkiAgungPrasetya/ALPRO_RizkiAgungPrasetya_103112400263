package main

import "fmt"

func main() {
	var bulan int
	var rumus int

	fmt.Println("Masukan tahun yang ingin di ketahui")
	fmt.Scanln(&bulan)
	rumus = bulan % 4 // logika apakah tahun input habis di bagi 4
	if rumus == 0 {   //jika pembagia tahun tidak tersisa maka kabisat
		fmt.Println("Tahun", bulan)
		fmt.Println("Kabisat : True")
	} else { //jika pembagian mempunyai nilai sisa maka bukan kabisat
		fmt.Println("Tahun", bulan)
		fmt.Println("Kabisat : False")
	}
}
