package main

import (
	"context"
	"log"
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
}
