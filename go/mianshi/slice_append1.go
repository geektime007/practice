package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := s[4:6]
	s[4] = 44
	s1 = append(s1, s[7])
	s[7] = 77
	fmt.Println("s=", s)
	s1 = append(s1, 100)
	s1 = append(s1, 101)
	s1 = append(s1, 102)
	s1 = append(s1, 103)
	fmt.Println("s=", s)
	fmt.Println(s1)
}
