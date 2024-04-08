package main

import (
	"fmt"
)

type z struct {
	int1, int2 int
}

type r struct {
	a []int
	b string
}

func main() {
	fmt.Println("Advanced New")

	var y *z = new(z)
	*y = z{89, 90}
	fmt.Println(y)

	var x *[2]r = new([2]r)
	x[0] = r{[]int{2, 3}, "hello"}
	x[1] = r{[]int{4, 5}, "world"}
	fmt.Println(x)
}
