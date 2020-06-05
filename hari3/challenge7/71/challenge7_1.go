package main

import "fmt"

func main() {
	var deret1 = []int{2, 7, 11, 15}
	var sum int = 9
	fmt.Println(sumPair(deret1, sum))
	deret1 = []int{3, 5, 2, -4, 8, 11}
	sum = 7
	fmt.Println(sumPair(deret1, sum))
}
func sumPair(arrayMasuk []int, total int) [][]int {
	var arrayBaru = [][]int{}
	for _, isi := range arrayMasuk {
		for i := 1; i < len(arrayMasuk); i++ {
			if isi+arrayMasuk[i] == total {
				arraySementara := []int{isi, arrayMasuk[i]}
				arrayBaru = append(arrayBaru, arraySementara)
				isi = 0
				arrayMasuk[i] = 0
			}
		}
	}
	return arrayBaru
}
