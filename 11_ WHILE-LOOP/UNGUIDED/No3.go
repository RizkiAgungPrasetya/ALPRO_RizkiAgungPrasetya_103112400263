package main

import (
	"fmt"
)

func main() {
	var bilbul, pembagi int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilbul)
	fmt.Print("Masukkan bilangan pembagi: ")
	fmt.Scan(&pembagi)

	jumlah := 0
	for bilbul >= pembagi {
		bilbul -= pembagi
		jumlah++
	}

	fmt.Print("Hasil div:", jumlah)
}
