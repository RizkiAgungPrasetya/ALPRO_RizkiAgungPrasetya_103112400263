package main

import "fmt"

func main() {
	var nomor int
	prima := true

	fmt.Scan(&nomor)

	for i := 2; i*i <= nomor; i++ {
		if nomor%i == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Print("Prima")
	} else {
		fmt.Print("Bukan Prima")
	}
}
