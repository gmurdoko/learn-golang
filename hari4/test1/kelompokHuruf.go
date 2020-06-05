// gays
// gw punya kumpulan Huruf "a", "b", "a", "u", "u", "c", "a"
// trus gw mau hitung jumlahnya per kelompok
// hasilnya, sbb:
// [map[a:3] map[b:1] map[u:2] map[c:1]]

// kalau mau coba pake potongan kode dibawah silahkan
// func main(){

// }
// func hitungVarian(h []string) []map[string]int {
//  // koding disini
//  }

//  func periksaMap(m []map[string]int, h string) (map[string]int, bool) {
//  // koding disini
//  }
package main

import (
	"fmt"
	"sort"
)

func main() {
	huruf := []string{"a", "b", "a", "u", "u", "c", "a"}
	arrayMapBaru := hitungVarian(huruf)
	fmt.Println(arrayMapBaru)
}
func hitungVarian(inputHuruf []string) []map[string]int {
	sort.Strings(inputHuruf)
	stringMap := []map[string]int{}
	sumHuruf := 0
	strSementara := ""
	for i, isi := range inputHuruf {
		if isi == inputHuruf[i] {
			sumHuruf++
			strSementara = isi
		} else {
			stringMap = append(stringMap, map[string]int{strSementara: sumHuruf})
			sumHuruf = 0
			strSementara = ""
		}

	}
	return stringMap

}
