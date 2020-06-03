package main

import (
	"fmt"
)

func main() {
	var (
		panjang     int
		lebar       int
		tinggi      int
		luasBalok   int
		volumeBalok int
	)
	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)
	luasBalok = panjang * lebar
	volumeBalok = panjang * lebar * tinggi
	fmt.Println("Luas permukaan: ", luasBalok)
	fmt.Println("Luas Volume: ", volumeBalok)
}
