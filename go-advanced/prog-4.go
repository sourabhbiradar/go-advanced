package main

import (
	"fmt"
)

func resume() {
	var name string = "Sourabh"
	var qual string = "BBA"
	var dob string = "15/07/1996"
	var address string = "#87 KSL Bluru-560010"
	var mobile int = 9087654321
	var email string = "email@gmail.com"
	var pan string = "AWD9898L"
	var skills string = "GoLang"
	var exp string = "1 yr as software engg."

	var info map[string]string = make(map[string]string)
	info["Name"] = name
	info["Qualification"] = qual
	info["DoB"] = dob
	info["Address"] = address
	info["email"] = email
	info["PAN"] = pan
	info["Skills"] = skills
	info["Experience"] = exp

	var cont map[string]int = make(map[string]int)
	cont["Contact"] = mobile
	fmt.Println(info)
	fmt.Println(cont)
}

func main() {
	resume()
}
