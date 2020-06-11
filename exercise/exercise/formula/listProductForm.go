package formula

import (
	// "exercise/main"
	"fmt"
)

func ListProductForm() {
	var main main.ProductList
	fmt.Println("Halaman Daftar Produk")
	fmt.Println("**************************************")
	fmt.Printf("%3s\t%-20s\t%-20s\n", "No.", "Kode Produk", "Nama Produk")
	for idx, product := range main.ProductList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", idx, product["productCode"], product["productName"])
	}
}
