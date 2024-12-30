package main

import "fmt"

func main() {

	var (
		uang, jumlah int
		selesai      bool
	)

	for selesai = false; !selesai; {

		fmt.Scan(&uang)

		jumlah += uang

		if uang <= 0 {
			selesai = true
		}
	}

	fmt.Print(jumlah)

}
