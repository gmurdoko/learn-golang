// pointer merupakan alamat memori
package main

import "fmt"

func main() {
	// angka := "20"
	// ubahAngka(&angka) //& mengirim alamat memori
	// fmt.Println(angka)
	// productList := make(map[string]int)
	// a := &productList
	// fmt.Printf("%p\n", &productList)
	// fmt.Println(&a)
	angka1 := 20
	var angka2 *int

	angka2 = &angka1
	*angka2 = 10
	fmt.Println(angka2, *angka2, &angka2)

}

// func ubahAngka(number *int) { //* untuk deklarasi alamat memori
// 	*number = 5 //mengumbah isi lamat
// }

//
