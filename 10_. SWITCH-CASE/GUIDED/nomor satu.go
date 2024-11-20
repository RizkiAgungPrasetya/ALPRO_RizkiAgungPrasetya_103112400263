package main

import "fmt"

func main() {
	var jam12, jam24 int
	var label string
	fmt.Print("Masukkan jam dalam format 24 jam: ")
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		jam12 = 12
		label = "AM"
	case jam24 < 12:
		jam12 = jam24
		label = "AM"
	case jam24 == 12:
		jam12 = 12
		label = "PM"
	case jam24 > 12 && jam24 < 24:
		jam12 = jam24 - 12
		label = "PM"
	default:
		fmt.Println("Tidak Termasuk Jam")
		return
	}

	fmt.Printf("Jam dalam format 12 jam: %d %s\n", jam12, label)
}
