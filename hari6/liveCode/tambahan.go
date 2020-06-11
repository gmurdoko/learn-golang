// 1. Jika angka tidak habis dibagi 2 atau 3, maka function akan
// me-return "SALAH".
// 2. Jika angka habis dibagi 2, maka return "TES"
// 3. Jika angka habis dibagi 3, maka return "LA"
// 4. Jika angka habis dibagi 2 dan 3, maka return "TESLA"

package main

import "fmt"

func main() {
	angka := 12
	hasil := ""
	teslaFunc(&angka, &hasil)
	fmt.Println(hasil)
}
func teslaFunc(i *int, hasil *string) {
	// hasil := ""
	if *i%2 == 0 && *i%3 == 0 {
		*hasil = "TESLA"
	} else if *i%2 == 0 {
		*hasil = "TES"
	} else if *i%3 == 0 {
		*hasil = "LA"
	} else {
		*hasil = "SALAH"
	}
	// return hasil
}
