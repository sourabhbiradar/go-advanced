package main

import (
	"fmt"
	"reflect"
)

func area2(z func(int, int) int) {

	fmt.Println("area2 :", z)
	fmt.Println(reflect.ValueOf(z))

}

func main() {
	ar := func(x, y int) int {
		area := 2 * (x + y)
		return area
	}

	fmt.Println(ar(3, 4))

	area2(ar)

}
