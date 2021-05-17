package main

import "fmt"

func test(s []int) {
	s = append(s, 4)
	fmt.Println("s:", s)
}

func test1(s *[]int) {
	*s = append(*s, 4)
	fmt.Println("s:", s)
}

func main() {
	fmt.Println("vim-go")
	s1 := make([]int, 0, 8)
	s1 = append(s1, 1)
	test(s1)
	fmt.Println("s1:", s1)
	test1(&s1)
	fmt.Println("s1:", s1)
}
