package main

import "fmt"

type isro struct {
	satalitte, rocket int
}

func (i *isro) New() {
	i.satalitte = 41
	i.rocket = 1
	fmt.Println(*i)
	var x *int = new(int)
	*x = 11
	//fmt.Println(*x)
	var y *isro = new(isro)
	y.satalitte = 32
	y.rocket = 108
	fmt.Println(*x, *y)

}
func main() {
	var chairman isro

	chairman.New()
}
