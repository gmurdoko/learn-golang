package formula

import "fmt"

func ListCategoryForm() {
	fmt.Println("Halaman Daftar Kategori")
	fmt.Println("**************************************")
	fmt.Printf("%3s\t%-20s\t%-20s\n", "No.", "Kode Kategori", "Nama Kategori")
	for idx, category := range categoryList {
		fmt.Printf("%-3d\t%-20s\t%-20s\n", idx, category["categoryCode"], category["categoryName"])
	}
}
