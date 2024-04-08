package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Simple Go Routine")
	go func() {
		fmt.Println("Go Routine in func main()")
	}()
	time.Sleep(3)

}
