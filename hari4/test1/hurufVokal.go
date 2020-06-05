package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "haha hihi hehe hoho huhu"
	hasilHitung := funHitungVokal(kata)
	fmt.Println("ini huruf vokalnya: ", hasilHitung)
}

func funHitungVokal(inputKata string) int {
	arrKataVokal := []string{"a", "i", "u", "e", "o"}
	sliceKata := strings.Split(inputKata, "")
	var kataVokal int
	for _, isiKata := range sliceKata {
		for _, isiVokal := range arrKataVokal {
			if isiKata == isiVokal {
				kataVokal++
			}
		}
	}
	return kataVokal
}
