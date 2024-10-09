package main

import (
	"fmt"
)

func main() {
	var a, b, c, bilangan int16

	fmt.Println("Masukan tiga digit")
	fmt.Scanln(&bilangan)

	a = bilangan / 100
	b = (bilangan % 100) / 10
	c = bilangan % 10

	fmt.Println(a <= b && b <= c)
}
