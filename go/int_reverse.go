package main

import "fmt"

func Reverse(in []int) []int {
	for a, b := 0, len(in)-1; a < b; a, b = a+1, b-1 {
		in[a], in[b] = in[b], in[a]
	}
	return in
}

func main() {
	fmt.Println("vim-go")
	in := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(in))
}
