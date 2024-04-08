package main

import "fmt"

func sadness() {
	var companyname string = "Google.inc"
	var salary int = 40000
	//var car string = "XUV 700"
	var car string = "nano"
	fmt.Println(companyname, car)
	if car != "XUV 700" {
		fmt.Println("Unhappy")
	} else if salary > 45000 {
		fmt.Println("Earn more")
	} else {
		fmt.Println("Happy")

	}
	if salary < 50000 {
		fmt.Println("sad")
	} else {
		fmt.Println("happy")
	}
}
func employees() {
	var employees [5]string = [5]string{"Akash,", "Yash,", "Kartik,", "Venkatesh,", "Abhi"}
	fmt.Println("Employees:", employees)
}
func emotions() {
	var person map[string]string = make(map[string]string)
	person["Yash"] = "happy"
	person["Abhi"] = "angry"
	person["Kartik"] = "smile"
	fmt.Println(person)
}
func main() {
	sadness()
	employees()
	emotions()
}
