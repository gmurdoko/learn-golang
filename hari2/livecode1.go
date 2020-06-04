package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputAngka int
	// var angkaPertama int
	// var angkaAhir int
	fmt.Print("Masukka angka: ")
	fmt.Scan(&inputAngka)
	// fmt.Printf("%T\n", inputAngka)
	nilaiAngka := strconv.Itoa(inputAngka)
	// fmt.Println(nilaiAngka)
	splitAngka := strings.Split(nilaiAngka, "")
	angkaPertama, _ := strconv.Atoi(splitAngka[0])
	angkaAhir, _ := strconv.Atoi(splitAngka[1])

	if angkaPertama%2 == 0 && angkaAhir%2 != 0 {
		fmt.Println("One Zombie Down!")
	} else if angkaPertama%2 == 0 && angkaAhir%2 == 0 {
		fmt.Println("Try Again to Attack")
	} else if angkaPertama%2 != 0 && angkaAhir%2 != 0 {
		fmt.Println("Try Again to Attack")
	} else if angkaPertama%2 != 0 && angkaAhir%2 == 0 {
		fmt.Println("You Are Dead!")
	}

}
