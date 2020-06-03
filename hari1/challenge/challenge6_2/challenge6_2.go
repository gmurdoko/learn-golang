package main

import (
	"fmt"
)

func main() {
	var angka int
	var ganjil int
	var genap int
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&angka)
	fmt.Print("Ganjil: ")
	for i := 1; i < angka; i++ {
		if i%2 != 0 {
			if i != 1 {
				fmt.Print(" + ")
			}
			fmt.Print(i)
			ganjil += i
		}
	}
	fmt.Println(" = ", ganjil)
	fmt.Print("Genap: ")
	for j := 1; j <= angka*2; j++ {
		if j%2 == 0 {
			if j != 2 {
				fmt.Print(" + ")
			}
			fmt.Print(j)
			genap += j
		}
	}
	fmt.Print("=", genap)
}
