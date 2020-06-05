package main

import (
	"strings"
)

func main() {
	var kata string = "kata kata kita mutiara semua"
	// fmt.Scan(&kata)
	// fmt.Println(kata)
	hasilFunKata := funHitungKata(kata)
	println(hasilFunKata)
}
func funHitungKata(inputKata string) int {
	spliceKata := strings.Split(inputKata, " ")
	hasil := len(spliceKata)
	return hasil
}
