package main

import "fmt"

func main() {
	// var text string

	// fmt.Print("Masukkan text: ")
	// fmt.Scan(&text)
	// for i := 0; i < len(text)/2; i++ {
	// 	if text[i] != text[len(text)-i-1] {
	// 		fmt.Println("false")
	// 	}
	// }
	// fmt.Println("true")

	var text, reverseText string
	fmt.Print("Masukkan text: ")
	fmt.Scan(&text)
	panjang := len(text) - 1
	for i := panjang; i >= 0; i-- {
		reverseText = reverseText + fmt.Sprint(string(text[i]))
	}
	if fmt.Sprint(text) == fmt.Sprint(reverseText) {
		fmt.Println("true")
	} else {
		fmt.Println(false)
	}
}
