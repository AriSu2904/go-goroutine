package main

import (
	"fmt"
	"time"
)

func main() {
	loopingNumber(5)
}

func loopingNumber(number int) {
	for i := 0; i < number; i++ {
		go displayNumber(i)
		fmt.Println("Inside Looping")
	}
	time.Sleep(1 * time.Second)
}

func displayNumber(num int) {
	fmt.Println(num)
}
