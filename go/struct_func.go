package main

import "fmt"

type A struct{}
type B struct{}
type C struct {
	A
	B
}

func (a *A) print() {
	fmt.Println("a")
}

func (b *B) print() {
	fmt.Println("b")
}

func main() {
	fmt.Println("vim-go")
	c := &C{}
	c.A.print()
	c.B.print()
	//c.print()
}
