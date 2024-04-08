package main

import "fmt"

type reserve interface {
	booking()
}
type movie struct {
	name   string
	price  int
	rating float32
}

func (p *movie) booking() {
	p.name = "BATMAN"
	p.price = 700
	p.rating = 8.9
	fmt.Println(*p)
}

type flight struct {
	from, destination string
	people            [4]string
}

func (f *flight) booking() {
	f.from = "Bluru"
	f.destination = "Mumbai"
	f.people = [4]string{"xyz,", "fgd,", "tyd,", "asw"}
	fmt.Println(*f)
}

type dinner struct {
	starter, main, dessert string
	items                  []int
	bill                   float32
}

func (d *dinner) booking() {
	d.starter = "paneer65"
	d.main = "biryani"
	d.dessert = "brownie"
	d.items = []int{2, 1, 2}
	d.bill = 897.80
	fmt.Println(*d)
}
func main() {
	var book reserve
	var mov movie
	book = &mov
	mov.booking()
	var flg flight
	book = &flg
	book.booking()
	var din dinner
	book = &din
	book.booking()
}
