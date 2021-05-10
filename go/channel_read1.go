package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var ch chan *int
	go func() {
		ch = make(chan *int, 1)
		a := 1
		ch <- &a
	}()

	go func(ch chan *int) {
		time.Sleep(time.Second)
		<-ch
		close(ch)
	}(ch)
	//go func() {
	//	time.Sleep(time.Second)
	//	<-ch
	//}()

	c := time.Tick(1 * time.Second)
	for range c {
		// 不传参数的方式一直打印1
		// 传参数的方式一直打印2
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
