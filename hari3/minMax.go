//min max array
package main

import (
	"fmt"
	"sort"
)

func main() {
	var deretInput = []int{1, 10, 45, 67, 33, 99, 2, 8}

	hasilMin, hasilMax, hasilAverage, hasilTotal := minMax(deretInput)
	fmt.Println("ini nilai min: ", hasilMin)
	fmt.Println("ini nilai max: ", hasilMax)
	fmt.Println("ini rata-rata: ", hasilAverage)
	fmt.Println("ini jumlah: ", hasilTotal)

}

func minMax(deretOlah []int) (int, int, float64, float64) {
	sort.Ints(deretOlah)
	max := deretOlah[len(deretOlah)-1]
	min := deretOlah[0]
	var total float64 = 0
	for _, value := range deretOlah {
		var newValue float64 = float64(value)
		total += newValue
	}
	average := total / float64(len(deretOlah))
	return min, max, average, total

}
