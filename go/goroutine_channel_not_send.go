package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("vim-go")
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		// 发送而不接收
		ch <- struct{}{}
	}()

	time.Sleep(time.Second)
}
