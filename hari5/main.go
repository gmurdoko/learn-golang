//masukkan kategori baru
//daftar kategori
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var productList []map[string]string
var categoryList []map[string]string

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

func mainMenu() {
	fmt.Println("**************************************")
	fmt.Println("Aplikasi")
	fmt.Println("**************************************")
	fmt.Println("1. Buat Produk Baru")
	fmt.Println("2. Daftar Produk")
	fmt.Println("3. Buat Kategori Baru")
	fmt.Println("4. Daftar Kategori")
	fmt.Println("5. Simpan File")
	fmt.Println("6. Keluar")
	fmt.Println("Pilihan menu dari 1-5")
}

func listProductForm() {
	fmt.Println("Halaman Daftar Produk")
	fmt.Println("**************************************")
	fmt.Printf("%3s\t%-20s\t%-20s\n", "No.", "Kode Produk", "Nama Produk")
	for idx, product := range productList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", idx, product["productCode"], product["productName"])
	}
}
func newProductForm() {
	var productCode string
	var productName string
	var saveProductConfirmation string

	fmt.Println("Halaman Buat Produk Baru")
	fmt.Println("**************************************")
	fmt.Print("Product Code : ")
	fmt.Scanln(&productCode)
	fmt.Print("Product Name : ")
	fmt.Scanln(&productName)
	fmt.Printf("Produk %s dengan kode %s akan disimpan (y/n) ? ", productName, productCode)
	fmt.Scanln(&saveProductConfirmation)

	if saveProductConfirmation == "y" {
		newProduct := make(map[string]string)
		newProduct["productCode"] = productCode
		newProduct["productName"] = productName
		productList = append(productList, newProduct)
		clearScreen()
	}
}
func newCategoryForm() {
	var categoryCode string
	var categoryName string
	var saveCategoryConfirmation string

	fmt.Println("Halaman Buat Kategori Baru")
	fmt.Println("**************************************")
	fmt.Print("category Code : ")
	fmt.Scanln(&categoryCode)
	fmt.Print("category Name : ")
	fmt.Scanln(&categoryName)
	fmt.Printf("Produk %s dengan kode %s akan disimpan (y/n) ? ", categoryName, categoryCode)
	fmt.Scanln(&saveCategoryConfirmation)

	if saveCategoryConfirmation == "y" {
		newCategory := make(map[string]string)
		newCategory["categoryCode"] = categoryCode
		newCategory["categoryName"] = categoryName
		categoryList = append(categoryList, newCategory)
		clearScreen()
	}
}
func listCategoryForm() {
	fmt.Println("Halaman Daftar Kategori")
	fmt.Println("**************************************")
	fmt.Printf("%3s\t%-20s\t%-20s\n", "No.", "Kode Kategori", "Nama Kategori")
	for idx, category := range categoryList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", idx, category["categoryCode"], category["categoryName"])
	}
}
func exportFileForm() {

}
func main() {
	mainMenu()
	for {
		var selectedMenu string

		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			newProductForm()
			break
		case "2":
			listProductForm()
			break
		case "3":
			newCategoryForm()
			break
		case "4":
			listCategoryForm()
			break
		case "5":
			exportFileForm()
			break
		case "6":
			os.Exit(0)
		default:
			clearScreen()
		}
	}
}
