package main

import (
	"fmt"
	"mahasiswa/model"
)

func main() {
	mahasiswa1 := model.NewMahasiswa("Sasa", "Matematika")
	mahasiswa2 := model.NewMahasiswa("Apin", "Kedokteran Binatang")
	fmt.Println(mahasiswa1.ToString())
	fmt.Println(mahasiswa2.ToString())
}
