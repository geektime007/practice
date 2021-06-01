package main

import (
	"fmt"
)

var a string

func main() {
	fmt.Println("vim-go")
	go func() {
		a = "Hello world"
	}()
	//time.Sleep(time.Second)
	fmt.Println(a)
}
