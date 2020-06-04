package main

import "fmt"

func main() {
	var name = "budi"
	age := 20
	// var name2 = "tono"
	greet(name, age)
	println(greet2(name))

	tambahUmur := umur(age)
	fmt.Println(tambahUmur)
}

func greet(namaCetak string, umurCetak int) {
	fmt.Printf("Selamat Datang %s umur %d\n", namaCetak, umurCetak)

}

func greet2(namaCetak string) string {
	// fmt.Printf("Selamat Datang %s\n", namaCetak)
	return namaCetak + " ini dari fungsi"
}
func umur(umurCetak int) int {
	return umurCetak + 20
}
