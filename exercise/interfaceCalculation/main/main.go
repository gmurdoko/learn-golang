package main

import (
	"exercise/interfaceCalculation/calcProcess"
	"fmt"
)

func main() {
	var input1, input2 int
	fmt.Print("Input angka 1 : ")
	fmt.Scan(&input1)
	fmt.Print("Input angka 2 : ")
	fmt.Scan(&input2)
	substraction := calcProcess.Pengurangan{input1, input2}
	substraction2 := calcProcess.Penambahan{input1, input2}
	printHasilHitung(substraction)
	printHasilHitung(substraction2)
}

func printHasilHitung(c calcProcess.Calculator) {
	hasil, pesanError := c.GetCalculation()
	if pesanError == nil {
		fmt.Println("Hasil Pengurangan : ", hasil)
	} else {
		fmt.Println(pesanError.Error())
	}

}
