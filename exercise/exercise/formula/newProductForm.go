package formula

import "fmt"

func NewProductForm() {
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
		ClearScreen()
	}
}
