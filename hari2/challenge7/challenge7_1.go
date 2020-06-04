package main

import "fmt"

func main() {
	var deret1 = []int{2, 7, 11, 15}
	var sum int = 9
	for i := 0; i < len(deret1); i++ {

		for j := i + 1; j < len(deret1); j++ {
			if deret1[i]+deret1[j] == sum {
				// fmt.Printf("[[%d %d]]\n", deret1[i], deret1[j])
				hasil := fmt.Sprintf("%d,%d", deret1[i], deret1[j])
				fmt.Printf("%s", hasil)

			}
		}
	}

	// var deret2 = []int{3, 5, 2, -4, 8, 11}
	// var sum2 int = 7
	// fmt.Print("[")
	// for a := len(deret2) - 1; a >= 0; a-- {

	// 	for b := len(deret2) - 1; b >= a; b-- {
	// 		// fmt.Println(deret2[b])
	// 		if deret2[a]+deret2[b] == sum2 {
	// 			fmt.Print("[", deret2[b], deret2[a], "]")
	// 		}
	// 	}

	// }
	// fmt.Print("]")
}
