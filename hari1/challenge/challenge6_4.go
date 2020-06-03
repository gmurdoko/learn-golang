package main

import "fmt"

func main() {
	var text string

	fmt.Print("Masukkan text: ")
	fmt.Scan(&text)
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			fmt.Println("false")
		}
	}
	fmt.Println("true")
}
