// interface bangunruang

// exp : struct persegi
//         struct lingkaran

// dicari : luas dan keliling

// dibuat interface
package main

import (
	"fmt"
	"math"
)

type bangunDatar interface {
	hitungLuas() float64
	hitungKeliling() float64
}

type persegi struct {
	panjang float64
}

type lingkaran struct {
	diameter float64
}

func main() {
	luas := persegi{15}
	printHasil(luas)

}

func printHasil(d bangunDatar) {
	fmt.Println("Luas Bangun Datar", d.hitungLuas())
	fmt.Println("Keliling Bangun Datar", d.hitungKeliling())
}

func (p persegi) hitungLuas() float64 {
	hasil := p.panjang * p.panjang
	// fmt.Println(hasil)
	return hasil
}
func (p persegi) hitungKeliling() float64 {
	hasil := 4 * p.panjang
	return hasil
}
func (p lingkaran) hitungLuas() float64 {
	hasil := math.Pi * math.Pow((p.diameter/2), 2)
	return hasil
}
func (p lingkaran) hitungKeliling() float64 {
	hasil := 2 * math.Pi * (p.diameter / 2)
	return hasil
}
