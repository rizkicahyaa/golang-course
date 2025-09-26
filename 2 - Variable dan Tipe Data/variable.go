package main

import "fmt"

func main() {
	// mendefinisikan tipe data
	var nama string = "Rizki"
	var umur int = 21

	// tidak mendefinisikan tipe data (otomatis terdeteksi)
	kota := "Yogyakarta"
	tahun := 2025

	fmt.Println("Nama: ", nama)
	fmt.Println("Umur: ", umur)
	fmt.Println("Kota: ", kota)
	fmt.Println("Tahun: ", tahun)
}