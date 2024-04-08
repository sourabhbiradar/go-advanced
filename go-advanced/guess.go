package main

import (
	"fmt"
	"math/rand"
	"os"
)


func main() {
	// guess the number

	var guess int
	fmt.Println("Enter a number between 1 to 10")
	fmt.Scanln(&guess)
	if guess < 0 || guess > 10 {
		fmt.Println("ENTER A NUMBER BETWEEN 1 to 10")
        os.Exit(0)
	}

	randomNum := rand.Intn(11) // 0 to n-1
    
    fmt.Printf("Your number was : %d\n",guess)
    fmt.Printf("Computer's number was : %d\n",randomNum)

	if guess == randomNum {
		fmt.Println("YOU WIN!!!")
	} else {
		fmt.Println("Better Luck Next Time")
	}

}
