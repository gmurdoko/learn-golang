package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var angka float64 = 30
	// age := 20

	// const firstname string = "murdoko"

	// fmt.Println(age)
	// fmt.Println(angka)
	// fmt.Println(firstname)

	// Operator perbandingan
	// var age int
	// var isEqual = (age == 18) //< > != == <= >= kondisi

	// fmt.Println(isEqual)

	//Operator Logika
	// var a = true
	// var b = false
	// var z = true

	// var c = a && b || z //operator logika AND
	// var d = a || b      //operator logika OR
	// var e = !b          //operator logika NEGASI / nilai kebalikan

	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	//Manipulasi String

	// convert

	// var a float64 = 3.9
	// var b int32 = int32(a)

	// fmt.Println(b)
	// fmt.Printf("%v\n", b)

	// tulisan := "halo"

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(tulisan)
	// }

	// if tulisan == "halo" {
	// 	fmt.Println("iya halo")
	// } else {
	// 	fmt.Println("bukan halo")
	// }

	// angka := 0
	// switch angka {
	// case 0:
	// 	fmt.Println("nol")
	// default:
	// 	fmt.Println("salah")

	// }
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Umur Anda: ")
	name, err := reader.ReadString('\n')
	fmt.Println(name)
	fmt.Println(err)

}
