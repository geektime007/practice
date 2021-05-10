package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(2)
	}()
	wg.Wait()
}
