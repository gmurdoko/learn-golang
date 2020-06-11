package operate

type trapesium struct {
	lebarAtas, lebarAlas, tinggi float64
}

func (isi trapesium) HitungLuas() float64 {
	return (isi.lebarAlas + isi.lebarAtas) * isi.tinggi / 2
}
func NewTrapesium(InputLebarAtas, InputLebarAlas, InputTinggi float64) *trapesium {
	return &trapesium{InputLebarAtas, InputLebarAlas, InputTinggi}

}
