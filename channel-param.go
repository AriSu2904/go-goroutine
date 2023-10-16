package main

import "fmt"

func main() {

	balance := make(chan int)
	go checkBalance(balance)
	myBalance := <-balance
	fmt.Println(myBalance) //100
}

func checkBalance(point chan int) {
	point <- 100
}
