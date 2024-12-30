package main

import "fmt"

func main() {

	var num, counter, d int

	fmt.Scan(&num)
	counter = 0

	for num > 0 {
		d = num % 10

		if d%2 == 0 && d != 0 {
			counter++
		}

		num /= 10

	}

	fmt.Println(counter)

}
