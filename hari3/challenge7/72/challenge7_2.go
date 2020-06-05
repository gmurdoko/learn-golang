package main

import "fmt"

func main() {
	deret1 := []int{6, 2, 4, 10, 5}
	// deret2 := []int{3, 9, 10, 13, 1, 2}
	var sementara int = 0
	for j := 0; j < len(deret1)-1; j++ {
		for i := 1; i <= len(deret1)-1; i++ {
			println(deret1[j], "\t", deret1[i])
			if deret1[j] > deret1[i] {
				sementara = deret1[j]
				deret1[j] = deret1[i]
				deret1[i] = sementara

			}

		}
	}
	fmt.Println(deret1)
}
