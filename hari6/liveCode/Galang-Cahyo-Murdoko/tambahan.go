// 1. Jika angka tidak habis dibagi 2 atau 3, maka function akan
// me-return "SALAH".
// 2. Jika angka habis dibagi 2, maka return "TES"
// 3. Jika angka habis dibagi 3, maka return "LA"
// 4. Jika angka habis dibagi 2 dan 3, maka return "TESLA"

package main

import "fmt"

func main() {
	angka := 7
	teslaFunc(angka)

}
func teslaFunc(i int) {
	// for i := 0; i < angka; i++ {
	if i%2 == 0 && i%3 == 0 {
		fmt.Println("TESLA")
	} else if i%2 == 0 {
		fmt.Println("TES")
	} else if i%3 == 0 {
		fmt.Println("LA")
	} else {
		fmt.Println("SALAH")
	}
	// }
}
