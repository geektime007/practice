package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("vim-go")
	go time.AfterFunc(3*time.Second, func() { fmt.Println("你好") })
	wg := sync.WaitGroup{}
	c := make(chan struct{})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeOut(&wg, time.Second*5) {
		close(c)
		fmt.Println("time out exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeOut(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回 false

	ch := make(chan bool, 1)
	go time.AfterFunc(timeout, func() { ch <- true })

	go func() {
		wg.Wait()
		ch <- false
	}()
	return <-ch
}
