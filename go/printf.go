package main

import "fmt"

func AA() int {
	a := 20
	return a
}

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%+v\n", AA())
	a := int(0)
	b := int(0)
	fmt.Println(a)
	fmt.Println(b)
}
