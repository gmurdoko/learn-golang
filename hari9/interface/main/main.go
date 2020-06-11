package main

import (
	"fmt"
	"hari9/interface/fauna"
)

func main() {
	kucingAnggora := fauna.NewKucing()
	duniaHewan := []fauna.Binatang{kucingAnggora}

	for _, fauna := range duniaHewan {
		cetak(fauna)
	}
}

func cetak(b fauna.Binatang) {
	fmt.Println(b.JumlahKaki())
}
