package main

import "fmt"

func main() {
	var jawaban, code int32

	jawaban = 8

	for {
		fmt.Print("Masukkan angka rahasia: ")
		fmt.Scan(&code)
		if code == jawaban {
			fmt.Print("Selamat jawaban anda benar!")
			break
		} else {
			fmt.Println("Tebakan anda salah, coba lagi.")
		}
	}
}
