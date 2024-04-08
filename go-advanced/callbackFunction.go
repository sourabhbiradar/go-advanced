package main

import (
	"fmt"
)

func movie(m string, callback func(int) int) {
	ret := callback(5)
	fmt.Println(ret)
}

func main() {
	fmt.Println("Call back function")
	movie("omg", func(folder int) int {
		return folder
	})

}
