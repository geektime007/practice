package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := "Hello world"
	for v := range a {
		fmt.Printf("%T %P %v\n", v, v, v)
	}
}
