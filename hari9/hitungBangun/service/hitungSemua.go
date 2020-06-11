package service

import "hari9/hitungBangun/operate"

type total struct {
	totalLuas []operate.BangunDatar
}

func (isi total) HitungLuas() float64 {
	var totalSemua float64
	for _, isinya := range isi.totalLuas {
		totalSemua += isinya.HitungLuas()
	}
	return totalSemua
}

func NewTotal(InputTotal []operate.BangunDatar) *total {
	return &total{InputTotal}

}
