package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	// Input nilai akhir mata kuliah
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Menentukan nilai mata kuliah berdasarkan nilai akhir
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	// Menampilkan hasil
	fmt.Println("Nilai mata kuliah: ", nmk)
}

//a. Jika nam diberikan 80.1, apa keluaran dari program tersebut?
//Jika nam adalah 80.1, program akan mengubahnya menjadi string ("A"), tetapi kemudian mencoba membandingkan string dengan angka, yang menyebabkan error. Ini tidak sesuai dengan spesifikasi soal.

//b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
// Tipe data tidak kompatibel: Program mencoba memberikan nilai string (misalnya "A") ke variabel nam yang bertipe float64, yang tidak dapat dilakukan.
//Pengecekan kondisi yang kurang tepat: Setelah nam diganti dengan string, kondisi berikutnya membandingkan antara string dan angka, yang akan menyebabkan error atau hasil yang tidak diinginkan.

//c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.
//Sudah Diperbaiki
