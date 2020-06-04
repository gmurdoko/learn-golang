package main

import (
	"fmt"
	"sort"
)

func olehOleh(saldo int, listHarga []int) {
	// var kosong int = 0
	sort.Ints(listHarga)
	// for i := 0; i < len(listHarga)-1; i++ {
	// 	for j := i + 1; j <= len(listHarga)-1; j++ {
	// 		if listHarga[i] > listHarga[j] {
	// 			kosong = listHarga[i]
	// 			listHarga[i] = listHarga[j]
	// 			listHarga[j] = kosong

	// 		}

	// 	}

	// }
	// fmt.Println(listHarga)
	var jumlahBarang int = 0
	for a := 0; a < len(listHarga); a++ {
		if listHarga[a] <= saldo {
			saldo -= listHarga[a]
			jumlahBarang++
		}

	}
	fmt.Println(jumlahBarang, "Barang")
	return
}
func main() {
	var saldo int = 30000
	var listHarga = []int{15000, 12000, 5000, 3000, 10000}
	olehOleh(saldo, listHarga)
	saldo = 10000
	listHarga = []int{2000, 2000, 3000, 1000, 2000, 10000}
	olehOleh(saldo, listHarga)
	// fmt.Println(olehOleh(30000, []int{15000, 12000, 5000, 3000, 10000}))
	// fmt.Println(olehOleh(10000, []int{2000, 2000, 3000, 1000, 2000, 10000}))
	// fmt.Println(olehOleh(4000, []int{7500, 1500, 2000, 3000}))
	// fmt.Println(olehOleh(50000, []int{25000, 25000, 10000, 15000}))
	// fmt.Println(olehOleh(0, []int{10000, 3000}))
}
