package main

import "fmt"

func main() {
	var input1, input2, input3, input4 string
	var a int

	fmt.Print("Masukan jawaban: ")

	a = 0

	for a <= 5 {
		fmt.Scan(&input1, &input2, &input3, &input4)
		if input1 == "merah" && input2 == "kuning" && input3 == "hijau" && input4 == "ungu" {
			fmt.Print("Berhasil: True")
			break
		} else {
			a++
			fmt.Print("Percobaan ", a, ": ", input1, input2, input3, input4, "\n")
		}

		if a == 5 {
			fmt.Print("Berhasil : False")
			break
		}
	}

}
