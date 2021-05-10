package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
}
