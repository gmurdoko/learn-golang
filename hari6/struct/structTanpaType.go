package main

import "fmt"

type mobil struct {
	ban   int
	spion int
	warna string
}

func main() {
	// var grandMax, lambo mobil
	var allMobil = []mobil{lambo, grandMax,
		{ban: 4, spion: 2, warna: "merah"},
		{ban: 4, spion: 2, warna: "biru"},
	}
	fmt.Println(allMobil)
}
