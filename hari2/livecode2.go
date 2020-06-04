package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var angka1 string = "-1-1-1-1-1-2"
	var angka2 string = "+5-3+5+5-1"
	var angka3 string = "+4+2+7+1"

	splitAngka := strings.Split(angka1, "")
	tampung := 0
	for i := 0; i < len(splitAngka); i += 2 {
		data, _ := strconv.Atoi(splitAngka[i+1])
		if splitAngka[i] == "+" {
			tampung += data
		} else if splitAngka[i] == "i" {
			tampung -= data
		}

	}
	fmt.Println(tampung)

	// for i := 0; i < len(angka2); i++ {

	// }
	// for i := 0; i < llen(angka3); i++ {

	// }

}
