package main

import (
	"fmt"
	"time"
)

func Produce(ch chan<- int, start int) {
	count := start
	for {
		ch <- count
		time.Sleep(time.Second)
		count++
	}
}

func Consumer(ch <-chan int) {
	for {
		select {
		case a, ok := <-ch:
			if !ok {
				fmt.Println("Consumer: over ")
				break
			}
			fmt.Println("Consumer: ", a)
		}
	}
}

func main() {
	fmt.Println("vim-go")
	ch := make(chan int, 10)
	go Consumer(ch)
	go Consumer(ch)
	go Produce(ch, 0)
	go Produce(ch, 1000)

	for {
		time.Sleep(10 * time.Second)
		close(ch) // bad method
	}
}
