package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string
	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)

	if strings.ContainsAny(email, "") == true {
		fmt.Println("email harus diisi, tidak boleh kosong")
	} else if strings.ContainsAny(email, "@") != true {
		fmt.Println("email harus mengandung @")
	} else if strings.ContainsAny(email, ".com") != true {
		fmt.Println("email tidak valid")
	} else if strings.ContainsAny(email, ".co.id") != true {
		fmt.Println("email tidak valid")
	} else if strings.ContainsAny(email, "*") == true {
		fmt.Println("email tidak boleh mengandung *")
	} else {
		fmt.Println("email benar")

	}

}
