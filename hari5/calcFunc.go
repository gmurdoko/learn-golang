// calc switch case
// operasi dengan function
// cetak function

package main

import "fmt"

func main() {
	angka1 := 4
	angka2 := 2
	operator := "*"
	switch operator {
	case "*":
		fmt.Println(perkalian(angka1, angka2))
	case "/":
		fmt.Println(pembagian(angka1, angka2))
	case "+":
		fmt.Println(penambahan(angka1, angka2))
	case "-":
		fmt.Println(pengurangan(angka1, angka2))
	default
		fmt.Println("operator tidak valid")
	}
}

var perkalian = func(angka1 int, angka2 int) int {
	hasil := angka1 * angka2
	return hasil
}

var pembagian = func(angka1 int, angka2 int) int {
	hasil := angka1 / angka2
	return hasil
}

var penambahan = func(angka1 int, angka2 int) int {
	hasil := angka1 + angka2
	return hasil
}
var pengurangan = func(angka1 int, angka2 int) int {
	hasil := angka1 - angka2
	return hasil
}

// func cetakHasil(hasil func(int, int) int, angka1, angka2 int) {
// 	fmt.Println(hasil(angka1, angka2))

// }
