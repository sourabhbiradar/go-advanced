package main

import (
	"fmt"
)

func main() {
	fmt.Println("Odd/Even in slice")
	var x []int = []int{}
	var y int
	for y = 0; y < 21; y++ {
		if y%2 != 0 {
			x = append(x, y)
			//fmt.Println(y)
		}
	}

	fmt.Println(x)

}
