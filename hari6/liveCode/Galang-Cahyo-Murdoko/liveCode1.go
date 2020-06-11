package main

import "fmt"

func main() {
	var koin = []int{1, 5, 7, 9, 11}
	jumlah := 11
	hasil := hitungKoin(koin, jumlah)
	fmt.Println(hasil)
}

func hitungKoin(koin []int, jumlah int) [][]int {
	var tampung = []int{}
	var hasil = [][]int{}
	for _, isi := range koin {
		// fmt.Println(isi)
		for i := 0; i < len(koin)-1; i++ {
			if isi+koin[i] == jumlah {
				tampung = append(tampung, isi, koin[i])
				hasil = append(hasil, tampung)
				isi = 0
				koin[i] = 0
				tampung = []int{}
				// return hasil

				break
			} else if isi+koin[i]+koin[i+1] == jumlah {
				tampung = append(tampung, isi, koin[i], koin[i+1])
				hasil = append(hasil, tampung)
				isi = 0
				koin[i] = 0
				tampung = []int{}

			}
		}

	}
	return hasil
}
