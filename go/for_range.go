package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("vim-go")
	m := make(map[string]*student)
	n := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 23},
		{Name: "li", Age: 24},
		{Name: "wang", Age: 21},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for i, _ := range stus {
		n[stus[i].Name] = &stus[i]
	}

	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", n)
}
