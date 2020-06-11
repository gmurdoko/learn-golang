package main

import (
	"exercise/interface/utils"
	"log"
)

func main() {
	// enDict := dictionary.EnglishDict{}
	// jpDict := dictionary.JapanDict{}

	// dictionaries := []dictionary.Dict{enDict, jpDict}

	// for _, dictionary := range dictionaries {
	// 	printGreeting(dictionary)
	// }
	result, err := pembagian(1, -1)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println(result)
	}

}

func pembagian(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, utils.DibagiNolError{}
	} else if num1 < 0 || num2 < 0 {
		return 0, utils.AngkaMinus{}
	}
	result := (num1 / num2)
	return result, nil

}

// func printGreeting(d dictionary.Dict) {
// 	fmt.Println(d.GetMorningGreeting())
// }
