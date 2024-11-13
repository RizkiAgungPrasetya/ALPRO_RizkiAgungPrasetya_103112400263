package main

import (
	"fmt"
	"unicode"
)

func main() {
	var char rune

	fmt.Print("Masukkan sebuah karakter: ")
	fmt.Scanf("%c", &char)

	if unicode.IsLetter(char) {
		// Jika karakter adalah huruf, kita periksa apakah itu vokal atau konsonan
		switch unicode.ToLower(char) {
		case 'a', 'e', 'i', 'o', 'u':
			fmt.Printf("%c adalah huruf vokal.\n", char)
		default:
			fmt.Printf("%c adalah huruf konsonan.\n", char)
		}
	} else {
		fmt.Printf("%c bukan merupakan huruf.\n", char)
	}
}