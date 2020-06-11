package main

import "fmt"

// var grandMax = mobil{ban: 1, spion: 2, warna: "Merah"}
// 	var toyota = mobil{1, 2, "Merah"}

// 	var banyakMobil = []mobil{grandMax, toyota, {ban: 2, spion: 1, warna: "Biru"}, {10, 2, "Jingga"}}
// 	fmt.Println(banyakMobil)

type mobil struct {
	merk  string
	ban   int
	spion int
	warna string
}

func main() {
	var allMobil = []mobil{
		{merk: "toyota", spion: 2, ban: 4, warna: "biru"},
		{ban: 2, spion: 2, warna: "merah", merk: "honda"},
		{ban: 6, spion: 4, merk: "suzuki", warna: "kuning"},
		{"bmw", 10, 8, "ijo"},
	}
	for _, isi := range allMobil {
		if isi.spion > 4 {
			fmt.Println(isi.merk)
		}
	}

}
