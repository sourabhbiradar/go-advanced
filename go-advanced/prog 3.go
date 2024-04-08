package main

import "fmt"

type library struct {
	book1, book2, book3 string
	pages               [3]int
	price               []float64
}

func (l *library) pass(b1 string, b2 string, b3 string, pg [3]int, prc []float64) {
	l.book1 = b1
	l.book2 = b2
	l.book3 = b3
	l.pages = pg
	l.price = prc
	fmt.Println(*l)
}
func (l *library) record(name string) {
	if l.book1 == "abc" && l.pages[0] < (1339) {
		fmt.Println(l.book1, name, "has been retured damaged\n", "please pay fine of rs", l.price[0])

	} else {
		fmt.Println(l.book1, name, "has been returned as it is")
	}

}
func (lbn *library) shift2() {
	if lbn.pages[1] != (1025) {
		fmt.Println("wrong book returned insted of", lbn.book2)
	} else {
		fmt.Println(lbn.book2, "retuned successfully")
	}
}
func school() {
	var sn string = "SGIS"
	fmt.Println("welcome to the libarary of", sn)

}
func main() {
	var info library
	school()
	info.pass("abc", "def", "ghi", [3]int{1334, 1027, 983}, []float64{349.9, 199.76, 400.43})
	info.record("Libarary Book")
	info.shift2()

}
