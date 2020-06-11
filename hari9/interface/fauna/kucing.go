package fauna

type kucing struct {
	kaki    int
	berbulu bool
}

func (k kucing) PunyaBulu() bool {
	k.berbulu = true
	return k.berbulu
}
func (k kucing) JumlahKaki() int {
	return k.kaki
}

func NewKucing() *kucing {
	return &kucing{4, true}
}
