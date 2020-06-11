package operate

type segitiga struct {
	lebar, tinggi float64
}

func (isi segitiga) HitungLuas() float64 {
	return isi.lebar * isi.tinggi / 2
}

func NewSegitiga(InputLebar, InputTinggi float64) *segitiga {
	return &segitiga{InputLebar, InputTinggi}

}
