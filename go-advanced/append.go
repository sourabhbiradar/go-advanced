package main

import "fmt"

func main() {
	fmt.Println("Append")
	var c []int = []int{1, 2, 3, 4}
	fmt.Println(c)
	c = append(c, 5, 6, 7) // sliceName=append(sliceName,ele1,ele2)
	fmt.Println(c)
}
