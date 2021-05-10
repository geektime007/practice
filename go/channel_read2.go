package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var ch chan int
	var count int

	go func() {
		ch <- 1
	}()

	go func() {
		count++
		time.Sleep(10 * time.Second)
		close(ch)
	}()

	<-ch

	fmt.Println(count)
}
