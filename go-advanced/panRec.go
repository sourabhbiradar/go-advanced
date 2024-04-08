package main

import "fmt"

func Panic() {
	s := "executing first()"
	fmt.Println(s)
	panic("error:panic")
	fmt.Println("done : first()") // not executed coz panic incurred before ths cud execute
}

func Recovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()
	fmt.Println("Recovery()")
	Panic()
	fmt.Println("done : Recovery()") // not executed coz of Panic() prog moved to defer func()
}
func main() {
	//Panic() //after panic no further program executes as Recovery() didnt execute--> i guess
	Recovery()
}
