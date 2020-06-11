package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "aku mencari kontain"
	cari := "aku"

	sliceKata := strings.Split(kata, " ")
	for indexKata, isiKata := range sliceKata {
		for i := indexKata; i <= len(sliceKata); i++ {
			if isiKata != cari {
				fmt.Println("tidak ada Kata-nya")
			} else {
				fmt.Println("ketemu Katanya")
			}

		}

	}

}
