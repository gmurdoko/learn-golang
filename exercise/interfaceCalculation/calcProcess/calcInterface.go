package calcProcess

type Calculator interface {
	GetCalculation() (int, error)
}
