package main

import "fmt"

type mobile struct {
	dvs1, dvs2, dvs3 string
}

func (m mobile) inti(m1 string, m2 string, m3 string) { //nt pointing to address bt instead copping the whole data
	m.dvs1 = m1
	m.dvs2 = m2
	m.dvs3 = m3
	fmt.Println(m, &m) //so whn '&m' executed its showing whole copy of 'm' nt jst memory address
}
func main() {
	var yash mobile
	yash.inti("apple", "samsung", "blackberry")
	fmt.Println(" BAD PROGRAMMING TECNIQUE\n")
}
