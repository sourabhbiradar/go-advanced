package main

import (
	"fmt"
)

func id(i int) int {

	if i == 0 {
		fmt.Println("in id()")
		return 1
	}
	fmt.Println(i)
	return (id(i - 1))

}

func main() {
	fmt.Println("Recurring function")

	ret := id(8)
	fmt.Println(ret)

}
