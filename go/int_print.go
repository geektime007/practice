package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := int(1)
	b := a - 128
	fmt.Printf("Type: %T, %v\n", b, b)

	c := int(1)
	d := c + 128
	fmt.Printf("Type: %T, %v\n", d, d)

	e := byte(1)
	f := e - 255
	fmt.Printf("Type: %T, %v\n", f, f)

	g := byte(1)
	h := g + 255
	fmt.Printf("Type: %T, %v\n", h, h)

}
