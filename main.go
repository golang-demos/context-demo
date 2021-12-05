package main

import (
	"context"
	"log"
	"strconv"
)

func main() {
	// Simple Background Context
	log.Print("Example #1: Simple Background context")
	ctx1 := context.Background()
	log.Print(ctx1)

	log.Print("")

	// Simple TODO Context
	log.Print("Example #2: Simple TODO context")
	ctx2 := context.TODO()
	log.Print(ctx2)

	log.Print("")

	// Simple Background Context with "myKey" holding value : 123
	log.Print("Example #3: Simple Background context with Value")
	ctx3 := context.WithValue(ctx1, "mykey", 123)
	log.Print(ctx3)
	log.Print("Value for the key \"mykey\":")
	log.Print(ctx3.Value("mykey"))

	log.Print("")

	log.Print("Example #4: Simple Background context with Cancel Function")
	ctx4, cancelFunc := context.WithCancel(ctx1)
	log.Print(ctx4)
	cancelFunc() // Canceling the context by calling cancel function
	log.Print(ctx4.Err())
	log.Print("Is canceled : " + strconv.FormatBool(ctx4.Err() == context.Canceled))
}
