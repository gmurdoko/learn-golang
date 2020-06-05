package main

import "fmt"

func main() {
	huruf := []string{"a", "b", "a", "u", "u", "c", "a"}
	hitungVarian(huruf)
}
func hitungVarian(inputHuruf []string) {
	var hurufMap = []map[string]int{}
	var hurufA = map[string]int{}
	var hurufB = map[string]int{}
	var hurufC = map[string]int{}
	var hurufU = map[string]int{}
	var angka int = 0
	for _, isi := range inputHuruf {
		if isi == "a" {
			angka++
			hurufA[isi] = angka
			angka = 0
		} else if isi == "b" {
			angka++
			hurufB[isi] = angka
			angka = 0
		} else if isi == "c" {
			angka++
			hurufC[isi] = angka
			angka = 0
		} else if isi == "u" {
			angka++
			hurufU[isi] = angka
			angka = 0
		}
	}
	hurufMap = append(hurufMap, hurufA, hurufB, hurufC, hurufU)
	fmt.Println(hurufMap)
}
