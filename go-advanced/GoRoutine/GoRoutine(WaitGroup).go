package main

import (
	"fmt"
	"sync"
)

func pm(s chan string) {
	fmt.Println("PM Modis statement :")

	fmt.Println(<-s)
}

func pak(s chan string) {
	fmt.Println("Pak statement :")
	fmt.Println(<-s)
}
func usa(o *sync.WaitGroup) {
	//fmt.Println("USAs statement :\n")
	fmt.Println("its just an accident ,India didnt do it on purpose.")

	o.Done() //o.Done()
}
func main() {
	fmt.Println("GoRoutine sync.WaitGroup\n")
	var statement chan string = make(chan string)
	go pm(statement)
	statement <- "sorry it was a mistake hehe."
	go pak(statement)
	statement <- "we will take it to UNSC."

	var obj sync.WaitGroup //sync.WAitGroup
	go usa(&obj)
	//fmt.Println("USAs statement :")

	obj.Add(1)//(1)=
	fmt.Println("Bidens statement :")
	obj.Wait()

}
