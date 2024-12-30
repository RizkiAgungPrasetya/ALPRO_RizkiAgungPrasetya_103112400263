package main

import "fmt"

func main() {

	var s, minggu int

	fmt.Scan(&s)

	minggu = s / 7

	if s%7 == 0 {
		fmt.Print("Minggu ke-", minggu)
	} else {
		fmt.Print("Minggu ke-", minggu+1)
	}

}
