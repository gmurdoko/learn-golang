package main

import (
	"hari9/hitungBangun/operate"
	"hari9/hitungBangun/print"
	"hari9/hitungBangun/service"
)

func main() {
	persegi1 := operate.NewPersegi(4, 2)
	persegi2 := operate.NewPersegi(4, 2)
	segitiga1 := operate.NewSegitiga(4, 2)
	trapesium1 := operate.NewTrapesium(10, 10, 11)

	ruangRuang := []operate.BangunDatar{persegi1, persegi2, segitiga1, trapesium1}
	totalSemua := service.NewTotal(ruangRuang)
	print.Cetak(persegi1)
	print.Cetak(persegi2)
	print.Cetak(segitiga1)
	print.Cetak(trapesium1)
	print.Cetak(totalSemua)

}
