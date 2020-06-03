package main

import (
	"fmt"
)

func main() {
	var angka int
	var ganjil int
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&angka)

	for i := 0; i < angka; i++ {
		if angka%2 == 0 {
			i++
			ganjil = ganjil + i
			fmt.Print(ganjil, "+")
		} else {
			angka = angka - 1
			i++
			ganjil = ganjil + i
			fmt.Print(ganjil, "+")
		}
	}
	fmt.Print("=")
	fmt.Println(ganjil)

	for i := 0; i < angka; i++ {
		if angka%2 != 0 {
			print(i)
		}
	}
}
