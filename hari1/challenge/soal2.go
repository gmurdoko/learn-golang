package main

import (
	"fmt"
)

func main() {
	alas := 20.0
	tinggi := 25.0

	//rumus luas segitiga = 1/2 alas * tinggi
	var hasil float64
	hasil = 0.5 * alas * tinggi

	fmt.Print("luas segitiga dengan alas= ", alas, " & tinggi= ", tinggi, " hasilnya =", hasil)

}
