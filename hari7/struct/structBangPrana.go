package main

import (
	"fmt"
)

type kunci string

type sliceKu []kunci

type addressMap map[kunci]sliceKu

type student struct {
	firstName, lastName kunci
	age                 int
	address
}

type address struct {
	homeAddress   addressMap
	officeAddress addressMap
}

func main() {
	student1 := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
		address: address{
			homeAddress: addressMap{"alamat1": []kunci{"subang"}, "alamat2": []kunci{"bandung"}},
		},
	}

	fmt.Println(student1)
}
