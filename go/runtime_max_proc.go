package main

import (
	"fmt"
	"runtime"
)

func main() {
	names := []string{"lily", "yoyo", "cersei", "rose", "annei"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
	}
	runtime.GOMAXPROCS(1)
	runtime.Gosched()
}
