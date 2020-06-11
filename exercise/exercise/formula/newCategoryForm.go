package formula

import "fmt"

func NewCategoryForm() {
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
		ClearScreen()
	}
}
