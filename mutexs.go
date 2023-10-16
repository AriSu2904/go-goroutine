package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var x = 0
	var mtx sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mtx.Lock()
				x += 1
				mtx.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(x) //100000
}
