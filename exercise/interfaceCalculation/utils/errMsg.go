package utils

type AngkaNegatifError struct {
	PesanMasalah string
}
type HasilKosong struct {
	PesanMasalah string
}

func (err AngkaNegatifError) Error() string {
	//	^Receiver			^fungsi ^return
	if len(err.PesanMasalah) == 0 {
		return MINUS
	}
	return err.PesanMasalah
}
func (err HasilKosong) Error() string {
	if len(err.PesanMasalah) == 0 {
		return OPERATION_ZERO
	}
	return err.PesanMasalah
}
