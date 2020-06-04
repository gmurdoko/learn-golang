package main

import "fmt"

func main() {
	// name := "budi"
	// age := 20
	// // var name2 = "tono"
	// greet(name, age)
	// println(greet2(name))

	// tambahUmur := umur(age)
	// fmt.Println(tambahUmur)

	// hasilLuas, hasilVolume := hitungSegitiga(10, 20)
	// fmt.Println(hasilLuas)
	// fmt.Println(hasilVolume)
	if ganjilGenap(1) {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
	// fmt.Println(ganjilGenap(1))
}

// func greet(namaCetak string, umurCetak int) {
// 	fmt.Printf("Selamat Datang %s umur %d\n", namaCetak, umurCetak)

// }

// func greet2(namaCetak string) string {
// 	// fmt.Printf("Selamat Datang %s\n", namaCetak)
// 	return namaCetak + " ini dari fungsi"
// }
// func umur(umurCetak int) int {
// 	return umurCetak + 20
// }

// func hitungSegitiga(alas, tinggi int) (luas, volume int) {
// 	luas = alas * tinggi / 2
// 	volume = alas * tinggi
// 	return luas, volume
// }

func ganjilGenap(input int) bool {
	if input%2 != 0 {
		return false
	} else {
		return true
	}
}
