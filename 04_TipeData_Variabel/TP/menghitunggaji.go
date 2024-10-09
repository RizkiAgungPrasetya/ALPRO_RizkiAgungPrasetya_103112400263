package main

import "fmt"

func main() {
	var hoursPerWeek float64
	var wagePerHour float64
	const normalHours = 40
	const overtimeMultiplier = 1.5

	// Meminta input jumlah jam kerja dalam seminggu dan upah per jam
	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scanln(&hoursPerWeek)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&wagePerHour)

	var normalPay, overtimePay float64

	// Jika jam kerja lebih dari 40 jam, hitung lembur
	if hoursPerWeek > normalHours {
		normalPay = normalHours * wagePerHour
		overtimeHours := hoursPerWeek - normalHours
		overtimePay = overtimeHours * wagePerHour * overtimeMultiplier
	} else {
		normalPay = hoursPerWeek * wagePerHour
		overtimePay = 0
	}

	// Menghitung total gaji mingguan
	totalWeeklySalary := normalPay + overtimePay

	// Menghitung gaji bulanan (4 minggu)
	totalMonthlySalary := totalWeeklySalary * 4

	// Menampilkan total gaji bulanan
	fmt.Printf("Total gaji bulanan: %.2f\n", totalMonthlySalary)
}
