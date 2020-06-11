package main

import "fmt"

type dict interface {
	GetMorningGreeting() string
}

type englishDict struct{}
type japanDict struct{}

func main() {
	enDict := englishDict{}
	jpDict := japanDict{}
	printGreeting(enDict)
	printGreeting(jpDict)
}

func printGreeting(d dict) {
	fmt.Println(d.GetMorningGreeting())
}

func (en englishDict) GetMorningGreeting() string {
	return "Hello, good morning"
}

func (jp japanDict) GetMorningGreeting() string {
	return "おはようございまず"
}
