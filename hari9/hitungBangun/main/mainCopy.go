package main

// import (
// 	"fmt"
// )

// type BangunDatar interface {
// 	HitungLuas() float64
// }
// type persegi struct {
// 	lebar  float64
// 	tinggi float64
// }

// func (isi persegi) HitungLuas() float64 {
// 	return isi.lebar * isi.tinggi
// }

// type segitiga struct {
// 	lebar, tinggi float64
// }

// func (isi segitiga) HitungLuas() float64 {
// 	return isi.lebar * isi.tinggi / 2
// }

// type trapesium struct {
// 	lebarAtas, lebarAlas, tinggi float64
// }

// func (isi trapesium) HitungLuas() float64 {
// 	return (isi.lebarAlas + isi.lebarAtas) * isi.tinggi / 2
// }

// type total struct {
// 	totalLuas []BangunDatar
// }

// func (isi total) HitungLuas() float64 {
// 	var totalSemua float64
// 	for _, isinya := range isi.totalLuas {
// 		totalSemua += isinya.HitungLuas()
// 	}
// 	return totalSemua
// }
// func main() {
// 	// persegi1
// 	// persegi2
// 	// segitiga
// 	// trapesium
// 	persegi1 := persegi{4, 2}
// 	tinggi1 := segitiga{4, 2}
// 	trapesium1 := trapesium{10, 10, 11}
// 	// cetak(trapesium1)
// 	ruangRuang := []BangunDatar{persegi1, tinggi1, trapesium1}
// 	totalSemua := total{ruangRuang}
// 	cetak(totalSemua)

// }
// func cetak(isi BangunDatar) {
// 	fmt.Println(isi.HitungLuas())
// }

// // func hitungLuasTotal(ruangRuang []BangunDatar) float64 {
// // 	totalLuas := 0.0
// // 	for _, isi := range ruangRuang {
// // 		totalLuas += isi.HitungLuas()
// // 	}
// // 	return totalLuas
// // }
