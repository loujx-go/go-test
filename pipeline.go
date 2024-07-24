package main

import "fmt"

func main() {
	for p := range buy3(buy2(buy(10))) {
		fmt.Println(p)
	}
}

func buy(n int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- fmt.Sprint("配件", i)
		}
	}()
	return ch
}

func buy2(in <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for c := range in {
			ch <- fmt.Sprint("组装（", c, ")")
		}
	}()
	return ch
}

func buy3(in <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for c := range in {
			ch <- fmt.Sprint("发货（", c, ")")
		}
	}()
	return ch
}
