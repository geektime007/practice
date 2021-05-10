package main

import (
	"fmt"
)

func fn() (n int) {
	defer func() (n int) {
		if v := recover(); v != nil {
			n, _ := v.(int)
			return n + 1
		}
		return
	}()
	n = 2
	panic(n)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(fn())
}
