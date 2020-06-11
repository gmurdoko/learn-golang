package print

import (
	"fmt"
	"hari9/hitungBangun/operate"
)

func Cetak(isi operate.BangunDatar) {
	fmt.Println(isi.HitungLuas())
}
