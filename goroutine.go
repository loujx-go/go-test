package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 10)

	go send(ch)

	//go receive()
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	defer close(ch)
	time.Sleep(3 * time.Second)
}

func send(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println(i)
	}
}
