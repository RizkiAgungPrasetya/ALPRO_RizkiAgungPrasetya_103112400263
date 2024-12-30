package main

import "fmt"

func main() {

	var bilangan float32

	fmt.Print("Masukkan Nilai TAK (1.0-4.0) : ")
	fmt.Scan(&bilangan)

	if bilangan >= 0.0 && bilangan <= 4.0 {
		if bilangan < 2.0 {
			fmt.Print("Poor")
		} else if bilangan >= 2.0 && bilangan <= 2.75 {
			fmt.Print("Fair")
		} else if bilangan > 2.75 && bilangan <= 3.0 {
			fmt.Print("Sastifactory")
		} else if bilangan > 3.0 && bilangan <= 3.5 {
			fmt.Print("Very Good")
		} else if bilangan > 3.5 {
			fmt.Print("Excellent")
		}
	} else {
		fmt.Println("Nilai TAK tidak sesuai")
	}

}
