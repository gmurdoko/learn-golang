package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputAngka int
	var angkaPertama int
	var angkaAhir int
	fmt.Print("Masukka angka: ")
	fmt.Scan(&inputAngka)
	fmt.Printf("%T\n", inputAngka)
	nilaiAngka := strconv.Itoa(inputAngka)
	fmt.Println(nilaiAngka)

	if angkaPertama% == 0 && angkaAhir% !=0{
		fmt.Println("One Zombie Down!")
	} else if angkaPertama% == 0 && angkaAhir% ==0 {
		fmt.Println("Try Again to Attack")
	} else if angkaPertama% != 0 && angkaAhir% !=0 {
		fmt.Println("Try Again to Attack")
	} else if angkaPertama% != 0 && angkaAhir% ==0 {
		fmt.Println("You Are Dead!")
	}

}
