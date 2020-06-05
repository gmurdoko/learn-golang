package main

import "fmt"

func main() {
	arrayRata := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ganjilGenap := "ganjil"
	var arrayGenap []int
	var arrayGanjil []int

	if ganjilGenap == "ganjil" || ganjilGenap == "genap" {
		for _, isi := range arrayRata {
			if isi%2 == 0 {
				arrayGenap = append(arrayGenap, isi)
			} else if isi%2 != 0 {
				arrayGanjil = append(arrayGanjil, isi)
			}
		}
		fmt.Println(arrayGanjil)
		fmt.Println(arrayGenap)
	} else {
		fmt.Println("Kondisi bukan ganjil atau genap")
	}

}
