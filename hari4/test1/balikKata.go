package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "bal"
	hasilKata := balikKata(kata)
	fmt.Printf("%s dibalik jadi %s", kata, hasilKata)

}

func balikKata(inputKata string) string {
	splitKata := strings.Split(inputKata, "")
	var balikKata string
	for i := len(splitKata) - 1; i >= 0; i-- {
		balikKata += splitKata[i]
	}
	return balikKata
}
