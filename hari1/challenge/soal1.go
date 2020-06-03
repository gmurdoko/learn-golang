package main

import (
	"fmt"
)

func main() {
	var a int = 15
	var b int = 20

	//kode disini

	a, b = b, a
	fmt.Print("a=", a, ", b=", b)

}
