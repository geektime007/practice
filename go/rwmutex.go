package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex

var count int

func main() {
	fmt.Println("vim-go")
	go A()
	time.Sleep(2 * time.Second)
	fmt.Println("Main Lock 1")
	//mu.Lock()
	//defer mu.Unlock()
	fmt.Println("Main Lock 2")
	count++
	fmt.Println(count)
	select {}
}

func A() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("A R Lock")
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("C R Lock")
}
