package main

import "fmt"

func main() {

	var bulan1, bulan2, profit float32

	fmt.Scan(&bulan1, &bulan2)

	profit = bulan2 - bulan1

	if profit > 0 {
		fmt.Println("Peningkatan Sebesar", profit)
	} else if profit == 0 {
		fmt.Println("Tetap")
	} else {
		fmt.Println("Penurunan Sebesar", -profit)
	}
}
