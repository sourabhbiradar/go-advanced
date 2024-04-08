package main

import (
	"fmt"
	"time"
)

func russia(b chan int) {
	fmt.Println(<-b, "bombs dropped")
	fmt.Println(len(b), cap(b))
}

func ukraine(i chan string) {

	fmt.Println(<-i)
	fmt.Println(len(i), cap(i))

}
func un() {

	fmt.Println("Russia attacked Ukraine!!")

}
func main() {
	fmt.Println("GoRoutine time\n")

	var bomb chan int = make(chan int)
	go russia(bomb)
	bomb <- 3

	go un()
	time.Sleep(0)

	var india chan string = make(chan string)
	go ukraine(india)
	india <- "we will talk with Putin to stop the war."

}
