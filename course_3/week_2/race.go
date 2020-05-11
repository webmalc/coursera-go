package main

import (
	"fmt"
	"time"
)

/*
Race Condition
A data race happens when two or more goroutines access
the same variable concurrently, and at least one of them
access is a write instruction.

To detect a race condition use the command:
go run -race race.go
*/

func race(i *int) {
	*i++
	fmt.Println(*i)
}

func main() {
	sharedCounter := 0
	go race(&sharedCounter)
	go race(&sharedCounter)
	time.Sleep(1 * time.Second)
}
