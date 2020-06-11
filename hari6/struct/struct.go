package main

import "fmt"

// type user struct {
// 	name string
// 	age  int
// }
type mobil struct {
	ban   int
	spion int
	warna string
}

func main() {
	var lambo mobil
	lambo.ban = 4
	lambo.spion = 2

	fmt.Println(lambo)
	// var user1 user
	// user1.name = "budi"
	// user1.age = 17
	// fmt.Printf("%T\n", user1)
	// fmt.Println(user1)
	// fmt.Println(user1.name)
	// fmt.Println(user1.age)

	// var user2 = user{"tono", 20}
	// fmt.Println(user2)

	// //cara laen tanpa type
	// var user3 = struct {
	// 	name string
	// 	age  int
	// }{"edo", 1}
	// fmt.Println(user3)

}
