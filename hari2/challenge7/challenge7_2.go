package main

import (
	"fmt"
	"sort"
)

func main() {
	deret1 := []int{6, 2, 4, 10, 5}
	deret2 := []int{3, 9, 10, 13, 1, 2}

	sort.Ints(deret1)
	fmt.Println("Terbesar: ", deret1[len(deret1)-1])
	fmt.Println("Terkecil: ", deret1[0])

	sort.Ints(deret2)
	fmt.Println("Terbesar: ", deret2[len(deret2)-1])
	fmt.Println("Terkecil: ", deret2[0])

}
