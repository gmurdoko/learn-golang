package main

import "fmt"

type student struct {
	firstName, lastName string
	age                 int
	studentAddress      address
}

type address struct {
	homeAddress map[string]string
	postCode    string
}

func main() {
	newStudent := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
		studentAddress: address{
			homeAddress: map[string]string{"alamat1": "subang", "alamat2": "bandung"},
			postCode:    "43183",
		},
	}

	fmt.Println(newStudent.studentAddress.homeAddress["alamat1"])
	// student1 := &newStudent
	// student2 := student1
	// student2.firstName = "ucap"
	// student2.lastName = "markicabs"
	// fmt.Println(newStudent)

	// student1 := student{
	// 	firstName: "budi",
	// 	lastName:  "anduk",
	// 	age:       18,
	// }

	// student2 := &student1
	// (*student2).firstName = "TONO"
	// fmt.Println(student2)

	// student2.firstName = "TINI"
	// fmt.Println(student2)

}
