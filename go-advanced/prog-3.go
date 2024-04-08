package main

import (
	"fmt"
)

func resume() {
	var name string = "Sourabh"
	var qual string = "BBA"
	var dob string = "15/07/1996"
	var address string = "#87 KSL Bluru-560010"
	//var mobile int = 9087654321
	var email string = "email@gmail.com"
	var pan string = "AWD9898L"
	var skills string = "GoLang"
	var exp string = "1 yr as software engg."

	var resume [8]string = [8]string{name, qual, dob, address, email, pan, skills, exp}
	fmt.Println("Resume : ", resume)

}
func main() {
	resume()
}
