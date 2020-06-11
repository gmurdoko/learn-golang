//finction bisa dijadikan parameter
//misal
package main

import "fmt"

//func as parameter
// var jumlah = func(a int, b int) int {
// 	return a + b
// }
// var kurang = func(a int, b int) int {
// 	return a - b
// }

func pertama() {
	fmt.Println("pertama")
}
func kedua() {
	fmt.Println("kedua")
}

func main() {

	defer kedua() //dieksekusi terakhir pada fungsi main
	pertama()
	// cetak(jumlah(1, 2)
	// a := 10
	// b := 5
	// c := 100
	//anonym parameter
	// func(penggali int) {
	// 	jumlah := (a + b) * penggali
	// 	fmt.Println(jumlah)
	// }(c)
	cetakBanyak("a", "b", "c")
}

// type hitung
//high ordered funtion

//variadic func
func cetakBanyak(s ...string) {
	for _, v := range s {
		fmt.Println("v")
	}
}

// func cetak(f func(int, int) int, a, b int) {
// 	fmt.Println(f(a, b))
// }

//deffered function
//pseudocode
