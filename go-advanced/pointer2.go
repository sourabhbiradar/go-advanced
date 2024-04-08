package main

import (
	"fmt"
)

type employee struct {
	emp1, emp2, emp3 string
	id1, id2, id3    int
}

func (p *employee) inti(e1 string, e2 string, e3 string, i1 int, i2 int, i3 int) {
	p.emp1 = e1
	p.emp2 = e2
	p.emp3 = e3
	p.id1 = i1
	p.id2 = i2
	p.id3 = i3
	fmt.Println(*p, &p)
}
func (p *employee) att() {
	if p.id1 == 420 {
		fmt.Println("Abbi is present")
	}
}
func main() {
	var JVtech employee
	JVtech.inti("Abhi,", "Yash,", "Kar,", 420, 555, 890)
	JVtech.att()
}
