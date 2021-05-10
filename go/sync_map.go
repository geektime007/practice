package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("vim-go")
	var m sync.Map
	// 没有m.Len()这个方法
	fmt.Println(m.Len())
	m.LoadOrStore("a", 1)
	m.Delete("a")
	fmt.Println(m.Len())
}
