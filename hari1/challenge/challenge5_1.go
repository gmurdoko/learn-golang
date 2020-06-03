package main

import (
	"fmt"
)

func main() {
	var nama string = "saitama"
	var peran string = "superhero"

	// kode disini

	if nama == "" && peran == "" {
		fmt.Println("Nama dan Peran harus diisi")
	} else if peran == "" {
		fmt.Println("Peran harus diisi")
	} else if peran == "superhero" {
		fmt.Println("Selamat datang Superhero Saitama, Kalahkan semua Monster di Bumi")
	} else if peran == "monster" {
		fmt.Println("Selamat datang Monster Saitama, Hancurkan semua Superhero yang Ada")
	} else {
		fmt.Println("Selamat datang Saitama, Pilih peranmu untuk melanjutkan game ini")
	}

}
