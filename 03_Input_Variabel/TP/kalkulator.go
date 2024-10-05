package main

import (
	"fmt"
)

func main() {

	var inp float32
	var ang1 float32
	var ang2 float32

	fmt.Println("Pilihlah Aritmatika berikut:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Pembagian")
	fmt.Println("4. Perkalian")
	fmt.Scanln(&inp)

	if inp == 1 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&ang1)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&ang2)
		fmt.Println("Hasil Penjumlahan adalah", ang1+ang2)
	} else if inp == 2 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&ang1)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&ang2)
		fmt.Println("Hasil Pengurangan adalah", ang1-ang2)
	} else if inp == 3 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&ang1)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&ang2)
		fmt.Println("Hasil pembagian adalah", ang1/ang2)
	} else if inp == 4 {
		fmt.Println("Masukan angka pertama")
		fmt.Scanln(&ang1)
		fmt.Println("Masukan angka kedua")
		fmt.Scanln(&ang2)
		fmt.Println("Hasil Perkalian adalah", ang1*ang2)
	} else {
		fmt.Print("Tidak ada pilihan bos yang bener aja")
	}
}
