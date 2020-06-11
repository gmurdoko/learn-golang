package model

import "fmt"

type mahasiswa struct {
	nama    string
	jurusan string
}

func (m *mahasiswa) ToString() string {
	sNama := fmt.Sprintf("%-15s %s\n", "Nama Lengkap", m.nama)
	sJurusan := fmt.Sprintf("%-15s %s\n", "Jurusan", m.jurusan)
	return sNama + sJurusan
}

//NewMahasiswa consturctor
func NewMahasiswa(nama, jurusan string) *mahasiswa {
	return &mahasiswa{
		nama:    nama,
		jurusan: jurusan,
	}
}
