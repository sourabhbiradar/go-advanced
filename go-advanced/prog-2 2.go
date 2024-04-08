package main

import "fmt"

type mobile struct {
	dvs1, dvs2, dvs3 string
}

func (m *mobile) inti(m1 string, m2 string, m3 string) {
	m.dvs1 = m1
	m.dvs2 = m2
	m.dvs3 = m3
	fmt.Println(*m, &m)
}
func (m *mobile) buy() {
	if m.dvs2 == "apple" {
		fmt.Println("best buy!")

	} else {
		fmt.Println("low budget")
	}

}
func main() {
	var yash mobile
	yash.inti("apple", "samsung", "blackberry")
	yash.buy()
}
