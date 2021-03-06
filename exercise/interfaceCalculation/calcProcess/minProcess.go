package calcProcess

import "exercise/interfaceCalculation/utils"

type Pengurangan struct {
	Angka1 int
	Angka2 int
}

func (p Pengurangan) GetCalculation() (int, error) {
	hasilNegatif := p.Angka1 - p.Angka2
	if hasilNegatif < 0 {
		return 0, utils.AngkaNegatifError{}
	} else {
		return hasilNegatif, nil
	}
}
