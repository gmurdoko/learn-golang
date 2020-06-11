package main

import "fmt"

func main() {
	for i := 0; i <= 25; i++ {
		isPrima := true
		if i == 0 || i == 1 {
			fmt.Printf("%d bukan bilangan prima\n", i)
			continue
		} else {
			for j := i - 1; j > 1; j-- {
				if i%j == 0 {
					fmt.Printf("%d bukan bilangan prima\n", i)
					isPrima = false
					break
				}
			}
			if isPrima {
				fmt.Printf("%d bilangan prima\n", i)
			}
		}

	}
}
