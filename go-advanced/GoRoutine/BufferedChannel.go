package main

import (
	"fmt"
	"time"
)

var c chan int = make(chan int)

func x(c chan int) {
	fmt.Println("help", <-c)
}
func y(c chan int) {
	fmt.Println("Police", <-c)
}
func main() {
	fmt.Println("Buffered Channel")
	var c chan int = make(chan int, 2)
	go y(c)
	go x(c)
	c <- 108
	c <- 100
	time.Sleep(1)
}
