//output slice
//input angka awal dan angka akhir
//angka awal < akhir
//ganjil genap

package main

import "fmt"

func main() {
	//inputan
	angkaPertama := 5
	angkaTerakhir := 10
	ganjilGenap := "ganjil"

	//Fungsi
	arrayFinal, status := fungsiArray(angkaPertama, angkaTerakhir)
	if status {
		fmt.Println("Kondisi benar array anda berisi", arrayFinal)
	} else {
		fmt.Println("Kondisi salah array tidak dicetak")
	}

	rataRataArray, statusRata := rataRata(arrayFinal)
	if statusRata {
		fmt.Println(rataRataArray, statusRata)
	} else {
		fmt.Println(rataRataArray, statusRata)
	}

	mainArrGanGen, statusGan := funGanjilGenap(arrayFinal, ganjilGenap)
	if statusGan {
		fmt.Println("Array: ", mainArrGanGen)
		fmt.Println(statusGan)
	} else {
		fmt.Println(statusGan)
	}
}
func fungsiArray(inputPertama, inputTerakhir int) ([]int, bool) {
	var array []int
	var statusFungsi bool
	if inputPertama > inputTerakhir {
		statusFungsi = false
	} else {
		for i := inputPertama; i <= inputTerakhir; i++ {
			array = append(array, i)
		}
		statusFungsi = true
	}
	return array, statusFungsi
}

func rataRata(arrayRata []int) (float64, bool) {
	//jumlah/banyak
	// var banyak float64 = float64(len(arrayRata))
	jumlah := 0
	if len(arrayRata) == 0 {
		return float64(jumlah), false
	}
	for _, isi := range arrayRata {
		jumlah += isi
	}
	// var jumlah2 float64 = float64(jumlah)
	rataRata := float64(jumlah / len(arrayRata))
	return rataRata, true

}

func funGanjilGenap(arrayRata []int, ganjilGenap string) ([]int, bool) {
	var arrayGanGen []int
	if ganjilGenap == "ganjil" {
		for _, isi := range arrayRata {
			if isi%2 == 0 {
				arrayGanGen = append(arrayGanGen, isi)
			}
		}
		return arrayGanGen, true
		// return arrayGanGen, arrayGanGen, status
	} else if ganjilGenap == "genap" {
		for _, isi := range arrayRata {
			if isi%2 != 0 {
				arrayGanGen = append(arrayGanGen, isi)
			}
		}
		return arrayGanGen, true

	} else {
		return arrayGanGen, false
	}

}
