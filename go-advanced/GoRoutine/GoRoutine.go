package main

import (
	"fmt"
)

func jeff(a chan string) {
	fmt.Println(<-a)

}
func main() {
	fmt.Println("Go Routine\n")
	var advocate chan string= make(chan string)
	go jeff(advocate)
	advocate <- " We lost furure bazar case against relaience group!!! "
}
