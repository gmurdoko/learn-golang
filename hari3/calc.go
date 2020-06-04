package main

import "fmt"

func main() {
	var inputAngka1, inputAngka2 float64
	var inputOperator string
	fmt.Print("Masukkin angka pertama: ")
	fmt.Scanln(&inputAngka1)
	fmt.Print("Masukkin angka kedua: ")
	fmt.Scanln(&inputAngka2)
	fmt.Print("Masukkin operator Hitungan: ")
	fmt.Scanln(&inputOperator)

	hasilHitung := kalkulaktor(inputAngka1, inputAngka2, inputOperator)
	fmt.Println("Hasil =", hasilHitung)
}

func kalkulaktor(angka1, angka2 float64, operator string) float64 {
	var hasil float64
	if operator == "*" {
		hasil = angka1 * angka2
	} else if operator == "/" {
		hasil = angka1 / angka2
	} else if operator == "+" {
		hasil = angka1 + angka2
	} else if operator == "-" {
		hasil = angka1 - angka2
	} else {

	}
	return hasil
}
