package operate

type persegi struct {
	lebar  float64
	tinggi float64
}

func (isi persegi) HitungLuas() float64 {
	return isi.lebar * isi.tinggi
}

func NewPersegi(InputLebar, InputTinggi float64) *persegi {
	return &persegi{InputLebar, InputTinggi}

}
