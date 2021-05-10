package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {

		}
	}()
	fmt.Println("Dropping mic")
	runtime.Gosched()
	runtime.GC() // 为啥现在不阻塞了
	fmt.Println("Done")
}
