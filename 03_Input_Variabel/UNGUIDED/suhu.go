package main

import (
	"fmt"
)

func main() {

	var f float32

	fmt.Println("Masukan suhu Celcius")
	fmt.Scanln(&f)
	fmt.Println("Jadi suhu Reamur nya adalah", f*4/5)
	fmt.Println("Jadi suhu Farenheit nya adalah", (f*9/5)+32)
	fmt.Println("Jadi suhu Kelvin nya adalah", f+273)
}
