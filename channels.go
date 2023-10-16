package main

import "time"

func main() {

	go func() {
		time.Sleep(5 * time.Second)
		hello <- "Hello From Anonymous Func"
	}()
	//result := <-hello
	//fmt.Println(result)
}

var hello = make(chan string)
