package main

import "fmt"

func Make() {
	//var x=
	//x:=
	var x []int = make([]int, 5) //can use 'make()' only for types slice,map and chan
	x[0] = 100
	x[1] = 99
	x[2] = 98
	x[3] = 97
	x[4] = 96
	fmt.Println(x)

	var y = make(map[string]string)
	y["a"] = "b"
	y["c"] = "d"
	y["e"] = "f"
	y["g"] = "h"
	y["i"] = "j"
	fmt.Println(y)

}
func z1(z chan float32) { //channel 'chan type' is only assosiated with Go Routine
	fmt.Println(<-z)
}
func main() {
	Make()
	z := make(chan float32)
	go z1(z)
	z <- 77.98

}
