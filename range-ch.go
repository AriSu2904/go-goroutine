package main

import (
	"fmt"
	"strconv"
)

func main() {

	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Perulangan Ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")

}
