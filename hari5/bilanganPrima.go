package main

import "fmt"

func main() {
	angka1 := 1
	angka2 := 30
	cekPrima(angka1, angka2)
}

func cekPrima(angka1, angka2 int) {
	for i := angka1; i <= angka2; i++ {
		if i > 1 {
			if i == 2 {
				fmt.Println(i)
			} else {
				for b := 2; b < i; b++ {
					if i%b == 0 {
						break
					} else if i-1 == b {
						fmt.Println(i)
					}
				}
			}
		}
	}

}
