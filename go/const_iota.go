package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota + 10
	d
)

const (
	name = "menglu"
	e    = iota
	f    = iota
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(a, b, c, d, e, f)
}
