package main

import (
	"fmt"
)

func main() {

	avgPoint := make(chan float32)
	totalPoint := make(chan int)

	go average(avgPoint, 70, 89, 90, 98, 68)
	go total(totalPoint, 70, 89, 90, 98, 68)

	for i := 0; i < 2; i++ {
		select {
		case avgResult := <-avgPoint:
			fmt.Println("Average Point :", avgResult)
		case totalResult := <-totalPoint:
			fmt.Println("Total Point :", totalResult)
		}
	}

}

func average(avg chan float32, numbers ...int) {
	var sum = 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	avg <- float32(sum) / float32(len(numbers))
}

func total(total chan int, numbers ...int) {
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	total <- sum
}
