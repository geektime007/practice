package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
