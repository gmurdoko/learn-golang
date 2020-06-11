package main

import (
	"fmt"
)

type bujurSangkar struct {
	sisi        float64
	luasnya     float64
	kelilingnya float64
}

func (b *bujurSangkar) keliling() {
	b.kelilingnya = 4 * b.sisi
}
func (b *bujurSangkar) luas() {
	b.luasnya = b.sisi * b.sisi
}

func main() {
	var bujurSaya = bujurSangkar{sisi: 10.0}
	bujurSaya.luas()
	bujurSaya.keliling()
	fmt.Println(bujurSaya.luasnya)
	fmt.Println(bujurSaya.kelilingnya)
}
