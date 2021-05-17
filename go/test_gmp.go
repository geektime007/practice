package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("vim-go")

	runtime.GOMAXPROCS(1)
	go func() {
		for {

		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("hello world")
}
