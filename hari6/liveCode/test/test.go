package main

import "fmt"

func main() {
	coin := []int{1, 5, 7, 9, 11}
	sum := 17

	if sum > 33 {
		fmt.Println("Kombinasi tidak ditemukan")
	} else {
		combi := findCoin(sum, coin)
		fmt.Println("Kombinasi yang dihasilkan adalah ", combi)
	}
}

func findCoin(total int, coins []int) [][]int {
	var result [][]int
	for i := 0; i < len(coins); i++ {
		for j := i + 1; j < len(coins); j++ {
			var check = coins[i] + coins[j]
			if check < total {
				for k := j + 1; k < len(coins); k++ {
					var checkAgain = coins[i] + coins[j] + coins[k]
					if checkAgain == total {
						var temp = []int{coins[i], coins[j], coins[k]}
						result = append(result, temp)
					}
				}
			}
			if check == total {
				var temp = []int{coins[i], coins[j]}
				result = append(result, temp)
			}
		}
	}
	return result
}
