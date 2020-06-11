package dictionary

//JapanDict is Model of Japan Dictionary
type JapanDict struct{}

//GetMorningGreeting is A Method Receiver
func (jp JapanDict) GetMorningGreeting() string {
	return "おはようございまず"
}
