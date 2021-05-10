package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}
