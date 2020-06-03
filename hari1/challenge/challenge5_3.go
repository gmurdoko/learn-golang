package main

import "fmt"

func main() {
	var angka int
	var hasil int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	hasil = angka % 2

	if hasil == 0 {
		fmt.Println("genap")
	} else {
		fmt.Println("ganjil")
	}

}
