//new Operator pointer
package main

import "fmt"

func main() {
	var alamatIntLangsung = new(int) //langsung alamat memori
	fmt.Println(*alamatIntLangsung)
	*alamatIntLangsung = 2
	fmt.Println(*alamatIntLangsung)
}
