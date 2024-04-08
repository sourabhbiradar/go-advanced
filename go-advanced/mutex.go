package main

import (
	"fmt"
	"sync"
)

var x = 100

func goR(msg chan string) {
	fmt.Println(<-msg, x)
	fmt.Println(x)

}
func mutex(m *sync.Mutex) {
	m.Lock()
	fmt.Println("Inside mutex()")
	var y = x - 1
	fmt.Println(y)
	m.Unlock()

}

func main() {
	fmt.Println("Mutex")

	var mtx sync.Mutex
	go mutex(&mtx)

	var msg = make(chan string)
	go goR(msg)
	msg <- "Go Routine executed"

}
