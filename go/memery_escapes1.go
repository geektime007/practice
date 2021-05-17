package main

import "fmt"

func test() {
	a := "abc"
	bs := []byte(a)
	fmt.Println(bs, len(bs), cap(bs))
}

func main() {
	fmt.Println("vim-go")
	test()
}
