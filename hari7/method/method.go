package main

import "fmt"

// method fungsi yang menempel pada type / struct

type student struct {
	firstName, lastName string
	age                 int
}

func (variabel student) fullName() string {
	return variabel.firstName + " " + variabel.lastName
}
func main() {
	newStudent := student{
		firstName: "Galang",
		lastName:  "Murdoko",
		age:       18,
	}

	namaLengkap := newStudent.fullName()
	fmt.Println(namaLengkap)

	// finalFullName := fullName(newStudent)
	// fmt.Println(finalFullName)
}

// func fullName(newStudent student) string {
// 	return newStudent.firstName + " " + newStudent.lastName

// }
