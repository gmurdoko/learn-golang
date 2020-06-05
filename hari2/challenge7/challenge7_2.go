package main

import (
	"fmt"
	"sort"
)

func main() {
	deret1 := []int{6, 2, 4, 10, 5}
	deret2 := []int{3, 9, 10, 13, 1, 2}

	iniMax, iniMin := maxMin(deret1)
	fmt.Println("Terbesar: ", iniMax)
	fmt.Println("Terkecil: ", iniMin)

	iniMax, iniMin = maxMin(deret2)
	fmt.Println("Terbesar: ", iniMax)
	fmt.Println("Terkecil: ", iniMin)

}

func maxMin(inputDeret []int) (max, min int, string) {
	sort.Ints(inputDeret)
	max = inputDeret[len(inputDeret)-1]
	min = inputDeret[0]
	return max, min
}
