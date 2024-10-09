package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

func main() {
	var a, b, c Point
	fmt.Scan(&a.X, &a.Y)
	fmt.Scan(&b.X, &b.Y)
	fmt.Scan(&c.X, &c.Y)

	ab := distance(a, b)
	ac := distance(a, c)
	bc := distance(b, c)

	fmt.Printf("Panjang sisi AB: %.2f\n", ab)
	fmt.Printf("Panjang sisi AC: %.2f\n", ac)
	fmt.Printf("Panjang sisi BC: %.2f\n", bc)

	var longestSide float64
	if ab >= ac && ab >= bc {
		longestSide = ab
	} else if ac >= ab && ac >= bc {
		longestSide = ac
	} else {
		longestSide = bc
	}

	fmt.Printf("Sisi terpanjang: %.2f\n", longestSide)
}
