package main

import (
	"context"
	"log"
	"strconv"
	"time"
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

	log.Print("")

	log.Print("Example #5: Simple Background context with Deadline specified in exact time")
	deadlineTimestamp := time.Now().Add(1 * time.Second) // Time after 1 second
	ctx5, _ := context.WithDeadline(ctx1, deadlineTimestamp)
	log.Print(ctx5)
	time.Sleep(2 * time.Second)
	log.Print(ctx5.Err())
	log.Print("Is canceled : " + strconv.FormatBool(ctx5.Err() == context.Canceled))
	log.Print("Is Deadline Exceeded : " + strconv.FormatBool(ctx5.Err() == context.DeadlineExceeded))

	log.Print("")

	log.Print("Example #6: Simple Background context with Deadline specified in exact time but canceled instead")
	deadlineTimestamp = time.Now().Add(1 * time.Second) // Time after 1 second
	ctx6, cancelFunc := context.WithDeadline(ctx1, deadlineTimestamp)
	log.Print(ctx6)
	deadline, ok := ctx6.Deadline()
	log.Print("Deadline Timestamp : " + deadline.String())
	log.Print("Is Deadline set? : " + strconv.FormatBool(ok))
	cancelFunc()
	time.Sleep(2 * time.Second)
	log.Print(ctx6.Err())
	log.Print("Is canceled : " + strconv.FormatBool(ctx6.Err() == context.Canceled))
	log.Print("Is Deadline Exceeded : " + strconv.FormatBool(ctx6.Err() == context.DeadlineExceeded))

}
