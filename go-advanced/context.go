package main

import (
	"context"
	"fmt"
	"time"
	"log"
)

func worker(ctx context.Context){
	select {
			case <-time.After(3* time.Second):
				fmt.Println("case: time.After()")
			case <-ctx.Done(): // received cancellation signal
				log.Println(ctx.Err().Error())
	}
}

func main() {
	// context
	fmt.Println("\tContext")

	fmt.Println("context.WithCancel()")

	ctx := context.Background() // creating parent/Root ctx
	ctx, cancel:= context.WithCancel(ctx)

    go worker(ctx) // worker goRoutine started

	time.Sleep(1* time.Second) // Waiting 3 sec before canceling the context
	cancel() // Cancel the context

	time.Sleep(1* time.Second) // Wait for the worker to finish
	fmt.Println("main executed")

}
