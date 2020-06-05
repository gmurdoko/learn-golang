package main

import (
	"fmt"
	"sort"
)

func main() {
	var saldo int = 30000
	var listHarga = []int{15000, 12000, 5000, 3000, 10000}
	totalBarang := olehOleh(saldo, listHarga)
	fmt.Println(totalBarang, "Barang")
	saldo = 10000
	listHarga = []int{2000, 2000, 3000, 1000, 2000, 10000}
	totalBarang = olehOleh(saldo, listHarga)
	fmt.Println(totalBarang, "Barang")
}

func olehOleh(saldo int, listHarga []int) int {
	sort.Ints(listHarga)
	var jumlahBarang int = 0
	for a := 0; a < len(listHarga); a++ {
		if listHarga[a] <= saldo {
			saldo -= listHarga[a]
			jumlahBarang++
		}
	}
	return jumlahBarang
}
