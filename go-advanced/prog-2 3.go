package main

import (
	"fmt"
)

type rectangle struct {
	length, breadth int
}
type circle struct {
	radius float64
}

func (r *rectangle) area1(l int, b int) {
	r.length = l
	r.breadth = b
	var areaR int = l * b
	fmt.Println("length =", l)
	fmt.Println("breadth =", b)
	fmt.Println("Area of Rectangle =", areaR)

}
func (c *circle) area2(r float64) {
	c.radius = r
	var Pi float64 = 3.14
	var areaC float64 = Pi * r * r
	fmt.Println("\nradius =", r)
	fmt.Println("Area of Circle =", areaC)

}

func main() {
	var sol1 rectangle
	sol1.area1(10, 12)
	var sol2 circle
	sol2.area2(14.6)

}
