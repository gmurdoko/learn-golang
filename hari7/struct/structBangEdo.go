package main

import "fmt"

// more advance gaysss
type arrayString []string
type addressMap map[string]arrayString // custom type
type student struct {
	firstName, lastName string
	age                 int
	address             // cara singkat ketika nama key sama dengan value
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
			homeAddress: addressMap{"alamat1": []string{"bandung"}, "alamat2": []string{"semanggi"}},
			// officeAddress: addressMap{"alamat1": "Bekasi", "alamat2": "Surabaya", "alamat3": "Semarang"},
		},
	}

	fmt.Println(student1.address.homeAddress["alamat1"][0])
}
