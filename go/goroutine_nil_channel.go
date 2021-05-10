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

	var ch chan int
	go func() {
		<-ch // 未初始化channel
	}()

	time.Sleep(time.Second)
}
