package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func test() {
	time.Sleep(time.Second * 2)
	os.Exit(-1)
}

func main() {
	fmt.Println("vim-go")
	runtime.GOMAXPROCS(1)
	go test()
	time.Sleep(10 * time.Second)

	for {
	}
}
