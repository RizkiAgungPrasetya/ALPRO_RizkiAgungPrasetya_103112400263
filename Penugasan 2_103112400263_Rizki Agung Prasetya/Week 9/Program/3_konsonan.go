package main

import "fmt"

func main() {

	// Variabel untuk menyimpan huruf
	var huruf string

	// Inputan untuk memasukan huruf
	fmt.Print("Masukan huruf: ")
	fmt.Scan(&huruf)

	// Percabangan untuk menentukan huruf vokal
	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" || huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf Vokal")
		//Ketika huruf vokal tidak terpenuhi, maka akan melakukan percabangan untuk huruf konsonan dan bukan huruf
	} else {
		// Percabangan untuk menentukan huruf konsonan
		if huruf >= "a" && huruf <= "z" || huruf >= "A" && huruf <= "Z" {
			fmt.Println("Huruf Konsonan")
			// Ketika huruf konsonan tidak terpenuhi, maka akan melakukan percabangan untuk bukan huruf
		} else {
			fmt.Println("Bukan Huruf")
		}
	}
}
