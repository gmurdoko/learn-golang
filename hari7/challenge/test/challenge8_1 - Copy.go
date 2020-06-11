// Bisa menampung sampai dengan 5 mahasiswa
// Validasi sesuai dengan screen, jika validasi tidak terpenuhi maka input tersebut akan terus diminta sampai terpenuhi
// Design screen harus mempunyai menu tambah mahasiswa, hapus mahasiswa, lihat mahasiswa dan exit
// Untuk menu tambah mahasiswa berikan validasi yaitu nama harus minimal 3 dan maksimal 20 karakter, umur minimal harus 17 tahun, jurusan maksimal 10 karakter
// Untuk menu hapus mahasiswa akan otomatis index terakhir yang terhapus
// untuk menu lihat mahasiswa akan ada sub menu lagi yaitu lihat menggunakan index dan lihat semua
// untuk menu lihat menggunakan index di input sesuai index data mahasiswa
// untuk menu lihat semua akan menampilkan seluruh data mahasiswa yang terdaftar
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var dbMahasiswa = []mahasiswa{}

type mahasiswa struct {
	nama, jurusan string
	umur          int
}

func main() {
	mainMenu()
	for {
		var selectMenu string

		fmt.Scanln(&selectMenu)
		switch selectMenu {
		case "1":
			addMahasiswa()
			break
		case "2":
			deleteMahasiswa()
			break
		case "3":
			viewMahasiswa()
			break
		case "4":
			os.Exit(0)
		default:
			clearScreen()
		}
	}

}

func mainMenu() {
	fmt.Println("--------------------------------------")
	fmt.Println("MAIN MENU")
	fmt.Println("--------------------------------------")
	fmt.Println("1. Add Mahasiswa")
	fmt.Println("2. Delete Mahasiswa")
	fmt.Println("3. View Mahasiswa")
	fmt.Println("4. Exit")
	fmt.Println("Pilihan menu")
}

func clearScreen() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	mainMenu()
}

func addMahasiswa() {
	var inputNama, inputJurusan, saveConfirm string
	var inputUmur int
	fmt.Println("--------------------------------------")
	fmt.Println("ADD MAHASISWA")
	fmt.Println("--------------------------------------")
	fmt.Print("Nama : ")
	fmt.Scanln(&inputNama)
	fmt.Print("Umur : ")
	fmt.Scanln(&inputUmur)
	fmt.Print("Jurusan : ")
	fmt.Scanln(&inputJurusan)

	if len(inputNama) >= 3 && len(inputNama) <= 20 && inputUmur >= 17 && len(inputJurusan) <= 10 {
		fmt.Print("Data Mahasiswa baru akan disimpan? (y/n) ")
		fmt.Scanln(&saveConfirm)
		if saveConfirm == "y" {
			mahasiswaBaru := mahasiswa{
				nama:    inputNama,
				umur:    inputUmur,
				jurusan: inputJurusan,
			}
			fmt.Println(mahasiswaBaru)
			dbMahasiswa = append(dbMahasiswa, mahasiswaBaru)
			fmt.Println(dbMahasiswa)
			clearScreen()
		} else {
			clearScreen()
		}

	} else {
		fmt.Println("\n Data yang anda masukkan tidak sesuai")
		// clearScreen()
		addMahasiswa()
	}

}

func viewMahasiswa() {
	var choose int
	fmt.Println("--------------------------------------")
	fmt.Println("VIEW MAHASISWA")
	fmt.Println("--------------------------------------")
	fmt.Println("1. Index Mahasiswa")
	fmt.Println("2. Tampilkan Semua")
	fmt.Println("Pilih menu : ")
	fmt.Scanln(&choose)

	if choose == 1 {
		var inputIndex int
		fmt.Print("Masukkan index mahasiswa : ")
		fmt.Scanln(&inputIndex)
		for index, isi := range dbMahasiswa {
			if index == inputIndex-1 {
				fmt.Println("--------------------------------------")
				fmt.Println("VIEW BY INDEX")
				fmt.Println("--------------------------------------")
				fmt.Println("Nama : ", isi.nama)
				fmt.Println("Umur : ", isi.umur)
				fmt.Println("Jurusan : ", isi.jurusan)
				fmt.Println("--------------------------------------")
			}
		}
		fmt.Println("Index kosong")

	} else if choose == 2 {
		fmt.Println("--------------------------------------")
		fmt.Println("VIEW ALL MAHASISWA")
		fmt.Println("--------------------------------------")
		for index, isi := range dbMahasiswa {
			fmt.Println(index + 1)
			fmt.Println("Nama : ", isi.nama)
			fmt.Println("Umur : ", isi.umur)
			fmt.Println("Jurusan : ", isi.jurusan)
			fmt.Println("--------------------------------------")

		}
	}

}

func deleteMahasiswa() {
	// var deleteConfirm string
	// fmt.Println("Apakah anda yakin akan menghapus data index terakhir? (y/n)")
	// fmt.Scanln(&deleteConfirm)
	// if deleteConfirm == "y" {
	dbMahasiswa[len(dbMahasiswa)-1] = mahasiswa{"", "", 0}
	dbMahasiswa = dbMahasiswa[:len(dbMahasiswa)-1]
	clearScreen()
	// }
}
