package main

import "fmt"

func main() {
	var ausername, apassword string

	username := "admin"
	password := "admin123"
	i := 0
	for {
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&ausername)
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&apassword)
		if ausername == username && apassword == password {
			fmt.Print(i, " Percobaan gagal login")
			break
		} else {
			i++
			fmt.Println("Username dan Password salah")
		}
	}
}
