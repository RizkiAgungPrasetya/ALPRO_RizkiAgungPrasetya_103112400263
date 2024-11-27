package main

import (
	"fmt"
	"strings"
)

func main() {
	const correctPassword = "#admin123" // Define the correct password
	const maxAttempts = 3
	attempts := 0

	for attempts < maxAttempts {
		fmt.Print("Masukkan password: ")
		var inputPassword string
		fmt.Scanln(&inputPassword)

		if strings.TrimSpace(inputPassword) == correctPassword {
			fmt.Println("Login berhasil!")
			return
		}

		attempts++
		remainingAttempts := maxAttempts - attempts

		if remainingAttempts > 0 {
			fmt.Printf("Password salah! Sisa kesempatan: %d\n", remainingAttempts)
		} else {
			fmt.Println("Login ditolak")
			return
		}
	}
}
