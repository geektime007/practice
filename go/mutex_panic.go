package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func main() {
	fmt.Println("vim-go")
	chain = "main"
	A()
	fmt.Println(chain)
}

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	// 这里会死锁
	//mu.Lock()
	//defer mu.Unlock()
	chain = chain + " --> C"
}
