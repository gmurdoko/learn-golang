package main

import "fmt"

type gaji struct {
	basic, tunjangan, lembur float64
}

type karyawan struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []gaji
}

func (e *karyawan) EmpInfo() string {

	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.basic)
		fmt.Println(info.tunjangan)
		fmt.Println(info.lembur)
		fmt.Println("Sub Total", info.basic+info.tunjangan+info.lembur)
	}
	return "----------------------"
}
func (e *karyawan) totalGaji() float64 {
	totalSemua := 0.0
	for _, isi := range e.MonthlySalary {
		totalSemua += isi.basic + isi.tunjangan + isi.lembur
	}
	return totalSemua

}

func main() {

	e := karyawan{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []gaji{
			gaji{
				basic:     15000.00,
				tunjangan: 5000.00,
				lembur:    2000.00,
			},
			gaji{
				basic:     16000.00,
				tunjangan: 5000.00,
				lembur:    2100.00,
			},
			gaji{
				basic:     17000.00,
				tunjangan: 5000.00,
				lembur:    2200.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())
	fmt.Println("Total : ", e.totalGaji())
}
