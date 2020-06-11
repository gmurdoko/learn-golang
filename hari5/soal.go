// Diberikan sebuah function pasanganTerbesar(angka)
//yang menerima 1 parameter berupa angka.
//Functiona akan menentukan pasangan dua digit angka mana yang paling besar dan me-return angka tersebut.
// Contoh, jika angka adalah 183928, maka function akan me-return 92,
// pasangan dua digit angka yang paling besar diantara yang lainnya.

package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	angka := 183928
	terbesar(angka)
}
func terbesar(angka int) {
	stringAngka := strconv.Itoa(angka)
	sliceString := strings.Split(stringAngka, "")
	stringPasang := ""
	var arrIntBaru = []int{}
	for i := 0; i < len(sliceString); i++ {
		stringPasang += sliceString[i]
		if len(stringPasang) == 2 {
			intAngka, _ := strconv.Atoi(stringPasang)
			arrIntBaru = append(arrIntBaru, intAngka)
			stringPasang = ""
		}
	}
	sort.Ints(arrIntBaru)
	println(arrIntBaru[len(arrIntBaru)-1])
}
