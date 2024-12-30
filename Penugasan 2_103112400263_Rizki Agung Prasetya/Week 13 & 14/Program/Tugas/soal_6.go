package main

import "fmt"

func main() {

	var masukan string
	var jumlahA, jumlahB, jumlahC int

	for {
		fmt.Scan(&masukan)

		if masukan == "A" {
			jumlahA++
		} else if masukan == "B" {
			jumlahB++
		} else if masukan == "C" {
			jumlahC++
		} else {
			break
		}
	}

	fmt.Println("Tipe A:", jumlahA)
	fmt.Println("Tipe B:", jumlahB)
	fmt.Println("Tipe C:", jumlahC)
}
