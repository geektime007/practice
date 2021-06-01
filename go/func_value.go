package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("vim-go")
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			fmt.Println("go1 i: ", i)
			wg.Done()
		}(&wg)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("go2 i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
