package main

import "fmt"

func main() {
	angka1 := 1
	angka2 := 5
	// fmt.Println(15 % 5)
	// fmt.Println(11 % 7)
	cariPrima(angka1, angka2)

}

func cariPrima(angka1, angka2 int) {

	for i := angka1; i <= angka2; i++ {
		b := 0
		for b = 2; b < i; b++ {
			// fmt.Println(i, b)
			if i%b == 0 {
				// fmt.Println(i)
				break
			} //else if i == b+1 {
			// fmt.Println(b)
			// }
		}
		// fmt.Println(i, b)
		if i == b {
			fmt.Println(i)
		}

	}
}
