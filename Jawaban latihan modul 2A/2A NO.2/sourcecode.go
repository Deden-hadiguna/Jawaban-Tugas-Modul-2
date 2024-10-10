package main

import (
	"fmt"
)

func main() {
	var tahun int

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	// Memeriksa tahun kabisat
	kabisat := false
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		kabisat = true
	}

	// Menampilkan hasil
	fmt.Println("Kabisat:", kabisat)
}