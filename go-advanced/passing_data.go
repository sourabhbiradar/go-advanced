package main

import "fmt"

type dt struct {
	name     string
	frnds    [5]string
	contacts []int
	//fans map[int]string
	//comm chan int

}

func (d *dt) pass(n string, f [5]string, c []int) {
	d.name = n
	d.frnds = f
	d.contacts = c
	fmt.Println(*d)
}

func main() {
	var fill dt
	fill.pass("sourabh", [5]string{"a", "b", "c", "d", "e"}, []int{900, 800, 700, 600, 500, 400, 300, 200})

}
