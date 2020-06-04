package main

import "fmt"

func main() {
	//Array
	//Golang kumpulan data yang bertipe sama
	//harus menentukan jumlah dulu
	// var kumpulanNama = [3]string{"antony", "garin", "ben"}
	// var kumpulanUmur [4]int
	//isi index
	// kumpulanUmur[0] = 10
	// kumpulanUmur[1] = 15
	// kumpulanUmur[2] = 20
	// fmt.Println(kumpulanNama)
	// fmt.Println(kumpulanUmur)
	// x := [...]int{10, 20, 30, 40} //tripledot
	// var y [5]int
	// z := [5]int{1: 10, 3: 30} //index
	// y = [5]int{10, 20, 30}

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	//Slice index yang tidak perlu
	// var kumpulanNama = []string{}
	// var kumpulanUmur = [...]int{10, 20, 30}

	// kumpulanNama = append(kumpulanNama, "galang", "cahyo")
	// kumpulanNama[0] = "hahaha"

	// fmt.Println(kumpulanUmur)
	// fmt.Println(kumpulanNama)

	// var kumpulanNama = []string{"budi", "tomo", "tini"} //jika slice alamat memori sama = data2 dirubah data satu ikut
	// var kumpulanNama2 = kumpulanNama                    //jika array tidak mereferensi alamat memori sama
	// kumpulanNama1 := make([]string, 1)
	//
	// kumpulanNama2[1] = "mahmut"

	// fmt.Println(kumpulanNama1)
	// fmt.Println(kumpulanNama)
	// fmt.Println(kumpulanNama2)

	//array ke slice
	// var kumpulan1 = [...]string{"a", "b", "c"}
	// var kumpulan2 = kumpulan1[:]
	// var kumpulan4 = kumpulan2[0:1]

	// fmt.Println(kumpulan4)
	// fmt.Println(len(kumpulan4)) // panjang array
	// fmt.Println(cap(kumpulan4)) // kapasitas array

	// kumpulan2 = append(kumpulan2, "kumpulan1")
	// fmt.Printf("%v\n", kumpulan2)

	//looping array
	// var kumpulan = [...]string{"budi", "tono", "tini", "sasa"}

	// for i := 0; i < len(kumpulan); i++ {
	// 	fmt.Println(kumpulan[i])
	// }
	// panjang := len(kumpulan)
	// fmt.Println(panjang)
	// fmt.Println(kumpulan)
	// for c := len(kumpulan) - 1; c >= 0; c-- {
	// 	// fmt.Println(c)
	// 	fmt.Println(kumpulan[c])
	// }

	//for range pada array
	var kumpulan = [...]string{"budi", "tono", "tini", "sasa"}
	for index, value := range kumpulan {
		fmt.Println(index, value)
	}
}
