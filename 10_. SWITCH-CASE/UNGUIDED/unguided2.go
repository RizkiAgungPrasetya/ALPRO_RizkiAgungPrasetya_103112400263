package main

import "fmt"

func main() {

	// Membuat variable
	var (
		kendaraan string
		durasi    int
	)

	// Meminta untuk memasukan jenis kendaraan dan durasi
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)

	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	// Switch case untuk menentukan tarif sesuai dengan jenis kendaraan dan durasi
	switch kendaraan {
	case "Motor", "motor", "MOTOR":
		switch {
		case durasi < 1:
			durasi = 1
			fmt.Print("Rp ", durasi*2000)
		default:
			durasi = durasi * 2000
			fmt.Print("Rp ", durasi)
		}
	case "Mobil", "mobil", "MOBIL":
		switch {
		case durasi < 1:
			durasi = 1
			fmt.Print("Rp ", durasi*5000)
		default:
			durasi = durasi * 5000
			fmt.Print("Rp ", durasi)
		}
	case "Truk", "truk", "TRUK":
		switch {
		case durasi < 1:
			durasi = 1
			fmt.Print("Rp ", durasi*8000)
		default:
			durasi = durasi * 8000
			fmt.Print("Rp ", durasi)
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}

}
