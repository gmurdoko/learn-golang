//masukkan kategori baru
//daftar kategori
package main

import (
	"exercise/formula"
	"fmt"
	"os"
)

type main struct {
	var ProductList []map[string]string
	var CategoryList []map[string]string
}



func exportFileForm() {

}
func main() {
	formula.MainMenu()
	for {
		var selectedMenu string

		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			formula.NewProductForm()
			break
		case "2":
			formula.ListProductForm()
			break
		case "3":
			formula.NewCategoryForm()
			break
		case "4":
			formula.ListCategoryForm()
			break
		case "5":
			exportFileForm()
			break
		case "6":
			os.Exit(0)
		default:
			formula.ClearScreen()
		}
	}
}
