package main

import "fmt"

func faktorial(n int) int {
	if n == 1 {
		return n
	} else {
		return n * faktorial(n-1)
	}
}

func main() {
	angka := faktorial(5)
	fmt.Println(angka)
}
