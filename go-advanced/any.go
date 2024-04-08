package main

import (
	"fmt"
)

func main() {
	// any
	fmt.Println("\ttype 'any'")
	// is Go really statically-typed lanuage ??

	var t any = 4
	fmt.Printf("t = %v ,%T\n", t, t)
	t = "t"
	fmt.Printf("t = %v ,%T\n", t, t)

}