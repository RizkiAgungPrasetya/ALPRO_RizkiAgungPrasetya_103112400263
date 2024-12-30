package main

import "fmt"

func main() {

	var gula, kopi, jumlahGula, jumlahKopi, terbuat int

	fmt.Scan(&gula, &kopi, &jumlahGula, &jumlahKopi)

	for kopi >= jumlahKopi && gula >= jumlahGula {
		kopi -= jumlahKopi
		gula -= jumlahGula

		terbuat++
	}

	fmt.Println(terbuat)

}
