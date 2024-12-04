package main

import "fmt"

func main() {
	var belanja, total int

	belanja = 0

	for i := 0; i < 1; {
		fmt.Print("Silahkan Masukan Harga Belanja Anda\njika ingin selelsai tekan 0: ")
		fmt.Scan(&total)
		if total == 0 {
			i++
		} else {
			belanja += total
		}
	}

	fmt.Print("Total belanja adalah: ", belanja)
}
