package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 通过发送而不接收

func queryAll() int {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func() { ch <- query() }()
	}
	// 有协程会被阻塞, 导致goroutine泄漏
	return <-ch
}

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func main() {
	fmt.Println("vim-go")
	//	fmt.Println("MaxGoroutine:", runtime.MaxGoroutine())
	for i := 0; i < 4; i++ {
		queryAll()
		fmt.Printf("goroutines: %d, CPU: %d\n", runtime.NumGoroutine(), runtime.NumCPU())
	}
}
