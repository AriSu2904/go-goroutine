package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string, 3)

	go func() {
		channel <- "John"
		channel <- "Doe"
		channel <- "Unknown"
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
	close(channel)
}
