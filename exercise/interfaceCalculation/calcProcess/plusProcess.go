package calcProcess

import "exercise/interfaceCalculation/utils"

type Penambahan struct {
	Angka1 int
	Angka2 int
}

func (p Penambahan) GetCalculation() (int, error) {
	hasilOperasi := p.Angka1 + p.Angka2
	if hasilOperasi == 0 {
		return 0, utils.HasilKosong{}
	} else {
		return hasilOperasi, nil
	}
}
