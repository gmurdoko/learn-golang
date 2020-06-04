package main

import "fmt"

func main() {
	//Hari Ke 3
	//Map
	// var warna = map[string]string{} //{"name": "budi"} //membuat map
	// //				^key   ^value

	// warna["merah"] = "warnaMerah"
	// warna["hijau"] = "warnaHijau"

	// var value , isExsist = warna["biru"]

	// if isExsist{
	// 	fmt.Println{
	// 		fmt.Println(value)
	// 	} else {
	// 		fmt.Println(value)
	// 	}
	// }

	// fmt.Println(warna)
	// fmt.Println(warna["alamat"])

	var warna = map[string]string{}

	warna["merah"] = "warnaMerah"
	warna["hijau"] = "warnaHijau"
	// for key, isi := range warna {
	// 	//^key ^isi
	// 	fmt.Println(key, isi)
	// }
	delete(warna, "merah") //delete map dengan key

	fmt.Println(warna)
}
