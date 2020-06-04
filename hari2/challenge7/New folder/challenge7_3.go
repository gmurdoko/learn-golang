package main

import (
	"fmt"
	"sort"
)

func olehOleh(saldo int, listHarga []int) {

	sort.Ints(listHarga)
	fmt.Println(listHarga)
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
	// var saldo int = 30000
	// var listHarga = []int{15000, 12000, 5000, 3000, 10000}
	// olehOleh(saldo, listHarga)
	// saldo = 10000
	// listHarga = []int{2000, 2000, 3000, 1000, 2000, 10000}
	// olehOleh(saldo, listHarga)
	olehOleh(30000, []int{15000, 12000, 5000, 3000, 10000})
	// olehOleh(10000, []int{2000, 2000, 3000, 1000, 2000, 10000})
	// olehOleh(4000, []int{7500, 1500, 2000, 3000})
	// olehOleh(50000, []int{25000, 25000, 10000, 15000})
	// olehOleh(0, []int{10000, 3000})
}
