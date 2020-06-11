// tepatnya type dan method receiver
// ini contoh
package main

import (
	"fmt"
	"strings"
)

type kata string

func (k kata) balikKata() string {
	var kataTerpisah = strings.Split(string(k), "")
	var hasil []string
	for i := len(kataTerpisah) - 1; i >= 0; i-- {
		hasil = append(hasil, kataTerpisah[i])
	}
	return strings.Join(hasil, "")
}

func (k kata) lowerCase() string {
	return strings.ToLower(string(k))
}

func (k kata) panjangKata() int {
	return len(k)
}

func main() {
	var kataSaya kata = kata("Edo")
	fmt.Println(kataSaya.lowerCase())
	fmt.Println(kataSaya.balikKata())
	fmt.Println(kataSaya.panjangKata())
}
