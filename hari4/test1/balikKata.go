package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "bal"
	hasilKata := balikKata(kata)
	fmt.Printf("%s dibalik jadi %s\n", kata, hasilKata)
	hasilKata2 := balikKata2(kata)
	fmt.Printf("%s dibalik jadi %s\n", kata, hasilKata2)
}

func balikKata(inputKata string) string {
	splitKata := strings.Split(inputKata, "")
	var balikKata string
	for i := len(splitKata) - 1; i >= 0; i-- {
		balikKata += splitKata[i]
	}
	return balikKata
}
func balikKata2(inputKata string) string {
	var balikKata string
	for i := len(inputKata) - 1; i >= 0; i-- {
		balikKata += inputKata[i : i+1]
	}
	return balikKata
}
