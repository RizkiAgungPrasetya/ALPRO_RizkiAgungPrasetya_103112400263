package main

import "fmt"

func main() {
	var code string

	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&code)
		if code == "TelU" {
			fmt.Print("Program selesai.")
			break
		} else {
			fmt.Println("Anda menegtik: ", code)
		}
	}
}
