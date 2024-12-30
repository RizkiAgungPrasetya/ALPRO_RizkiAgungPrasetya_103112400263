package main

import "fmt"

func main() {

	var h1, h2, m1, m2, mulai, selesai, durasi int

	fmt.Scan(&h1, &m1, &h2, &m2)

	mulai = h1*60 + m1
	selesai = h2*60 + m2

	if selesai < mulai {
		selesai += 12 * 60
	}

	durasi = selesai - mulai

	fmt.Println(durasi/60, "jam", durasi%60, "menit")

}
