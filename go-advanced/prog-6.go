package main

import (
	"fmt"
)

type list struct {
	item1, item2 string
	qnt          int
}

func (p *list) inti(e string, m string, q int) {
	p.item1 = e
	p.item2 = m
	p.qnt = q
	fmt.Println(*p, &p)
}
func (p *list) gros() {                                       //Why ths func ?
	var gros map[string]int = make(map[string]int)
	//gros[p.item1] = p.qnt                                     
	gros[p.item1]=q
	//gros[p.item2] = p.qnt                                    
	 gros[p.item2]=q
	fmt.Println(gros)
}
func main() {
	var glist list
	glist.inti("Eggs", "Milk", 6)
	glist.gros()
}
