package main

import (
	"fmt"
)

func resume() {
	var name string = "Sourabh\n"
	var qual string = "BBA\n"
	var dob string = "15/07/1996\n"
	var address string = "#87 KSL Bluru-560010\n"
	var mobile int = 9087654321
	var email string = "email@gmail.com\n"
	var pan string = "AWD9898L\n"
	var skills string = "GoLang\n"
	var exp string = "1 yr as software engg."

	fmt.Println(" Name : ", name, "Qualifications : ", qual, "Date of Birth : ", dob, "Address : ", address, "Contact : ", mobile, "\n eMail : ", email, "PAN : ", pan, "Skills : ", skills, "Experience : ", exp)
}
func main() {
	resume()
}
