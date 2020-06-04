package main

import "fmt"

func main() {
	//cetak berdasarkan nik
	//cetak semua
	var karyawan = map[int]string{}
	karyawan[1] = "Budi"
	karyawan[2] = "Makmur"
	karyawan[3] = "Tono"

	for nik, namaKaryawan := range karyawan {
		fmt.Println(nik, "\t", namaKaryawan)
	}

	var inputNik int
	fmt.Print("Masukkan nik yang akan dicari:")
	fmt.Scanln(&inputNik)

	for nik, namaKaryawan := range karyawan {
		if nik == inputNik {
			fmt.Println(nik, namaKaryawan)
		} else {
			fmt.Println("orangnya kaga ada")
		}
	}

}
